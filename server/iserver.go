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
)

type IServer struct{
	Config 				utils.Config
	DataServerClinet 	client.XClient
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

func (self *IServer)ConnectDataServer() client.XClient{
	zkAddr := self.Config.GetZookeeperIp()
	zd := client.NewZookeeperDiscovery("/"+self.Config.BasePath, "DataServer", zkAddr, nil)
	self.DataServerClinet = client.NewXClient("DataServer", client.Failover, client.RandomSelect,zd, client.DefaultOption)
	self.DataServerClinet.Auth(self.Config.ServerToken)
	return self.DataServerClinet
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