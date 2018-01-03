package iserver

import (
	"context"
	"errors"
	"time"
	"../log"
	"../utils"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/serverplugin"
	metrics "github.com/rcrowley/go-metrics"
)

type IServer struct{
	Config utils.Config
}

func (self *IServer)auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == self.Config.ServerToken {
		log.Log(log.Info,"Authentication Success")
		return nil
	}
	log.Log(log.Info,"Authentication Fail")
	return errors.New("invalid token")
}

func addRegistryPlugin(s *server.Server,basePath string,addr string,zkAddr []string) {
	zookeeperPlugin := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ZooKeeperServers: zkAddr,
		BasePath: basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	} 
	err := zookeeperPlugin.Start()

	if err != nil {
		log.Log(log.Fatel,"%s",err)
	} 
	s.Plugins.Add(zookeeperPlugin)
}

func (self *IServer) Init(){

}

func (self *IServer)ConnectDataServer(){

}

func (self *IServer) Start(name string,rcvr interface{}) {

	var err error
	self.Config,err = utils.ReadConfig()

	if(err == nil){
		zkAddr := self.Config.GetZookeeperIp()
		addr := self.Config.GetServerIp(name)
		s := server.NewServer()
		addRegistryPlugin(s,self.Config.BasePath, addr,zkAddr)
		s.RegisterName(name, rcvr, "")
		s.AuthFunc = self.auth
		s.Serve("tcp", addr)
	}
} 