package gateserver

import (
	"../../log"
	"../../utils"
	"github.com/smallnest/rpcx/client"
)

func Start() {

	config,err := utils.ReadConfig()

	if( err == nil){
		zkAddr := config.GetZookeeperIp()
		gw := NewGateway(config.BasePath,config.ServerToken,config.GateServer.Ip,zkAddr,  ServerType("http1"), client.Failtry, client.RandomSelect, client.DefaultOption)
		log.Log(log.Info,"start gate server")
		gw.Serve()
	}
}