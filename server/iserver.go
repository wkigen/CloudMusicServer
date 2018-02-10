package iserver

import (
	"context"
	"errors"
	"time"
	"../utils"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/serverplugin"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/golang/glog"
	"github.com/garyburd/redigo/redis"
)

type IServer struct{
	Config 				utils.Config
	DataServerXC 		client.XClient
	RedisPool 			redis.Pool
}

func (self *IServer)auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == self.Config.ServerToken {
		return nil
	}
	glog.Infoln("Authentication Fail")
	return errors.New("invalid token")
}

func addRegistryPlugin(s *server.Server,basePath string,addr string,zkAddr []string) {
	zookeeperPlugin := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ZooKeeperServers: zkAddr,
		BasePath: "/"+basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	} 
	err := zookeeperPlugin.Start()

	if err != nil {
		glog.Fatal("%s",err)
	} 
	s.Plugins.Add(zookeeperPlugin)
}

func (self *IServer)ConnectDataServer() {
	zkAddr := self.Config.GetZookeeperIp()
	zd := client.NewZookeeperDiscovery("/"+self.Config.BasePath, "DataServer", zkAddr, nil)
	self.DataServerXC = client.NewXClient("DataServer", client.Failover, client.RandomSelect,zd, client.DefaultOption)
	self.DataServerXC.Auth(self.Config.ServerToken)
}

func (self *IServer)ConnectRedis() {
	redisIp := self.Config.Redis.Ip
	password := self.Config.Redis.Password

	self.RedisPool = redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisIp)
			if err != nil {
				glog.Errorln("can not connect reids ip:",redisIp)
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				glog.Errorln("can not auth reids password:",password)
				c.Close()
				glog.Infoln(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (self *IServer) SetRedis(name string,key string,exTime int) (reply interface{}, err error){
	rds := self.RedisPool.Get()
	defer rds.Close() 
	return rds.Do("SET",name,key, "EX", exTime)
}

func (self *IServer) GetRedis(name string) (string,  error){
	rds := self.RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("GET",name))
}

func (self *IServer) Init() error{
	var err error
	self.Config,err = utils.ReadConfig()
	return err
}

func (self *IServer) Start(name string,rcvr interface{}) {

	zkAddr := self.Config.GetZookeeperIp()
	ip := self.Config.GetServerIp(name)
	port := self.Config.GetServerPort(name)
	s := server.NewServer()
	addRegistryPlugin(s,self.Config.BasePath, ip+":"+port,zkAddr)
	s.RegisterName(name, rcvr, "")
	s.AuthFunc = self.auth
	s.Serve("tcp", ":"+port)
} 

func (self *IServer) Stop() {
	glog.Flush()
}