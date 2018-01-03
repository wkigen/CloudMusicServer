package iserver

import (
	"context"
	"errors"
	"time"
	"../log"
	"../common"
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/serverplugin"
	metrics "github.com/rcrowley/go-metrics"
)

var (
	zkAddr = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/cloudmusic", "prefix path")
)

type IServer struct{
}

func (self *IServer)auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == common.ServerToken {
		log.Log(log.Info,"Authentication Success")
		return nil
	}
	log.Log(log.Info,"Authentication Fail")
	return errors.New("invalid token")
}

func addRegistryPlugin(s *server.Server,addr string) {
	zookeeperPlugin := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath: *basePath,
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

func (self *IServer) Start(name string,rcvr interface{},addr string ) {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s,addr)
	s.RegisterName(name, rcvr, "")
	s.AuthFunc = self.auth
	s.Serve("tcp", addr)
} 