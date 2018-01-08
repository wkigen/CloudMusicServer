package gateserver

import (
	"../../utils"
	"github.com/golang/glog"
	"github.com/smallnest/rpcx/client"
)

func Start() {
	glog.Infoln("start gate server")
	config,err := utils.ReadConfig()

	if( err == nil){
		zkAddr := config.GetZookeeperIp()
		gw := NewGateway(config.BasePath,config.ServerToken,config.GateServer.Ip,zkAddr,  ServerType("http1"), client.Failtry, client.RandomSelect, client.DefaultOption)
		gw.Serve()
	}
}