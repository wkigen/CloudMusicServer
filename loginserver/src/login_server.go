package loginserver


import (
	"../../server"
	"github.com/smallnest/rpcx/client"
	"github.com/golang/glog"
)

type ServerEntity struct {
	iserver.IServer
}

type LoginServer struct{
	DataServerXC client.XClient
}

var g_Entity ServerEntity

func Start(){
	glog.Infoln("LoginServer is start")
	loginServer := LoginServer{}
	g_Entity = ServerEntity{}
	err := g_Entity.Init()
	if(err != nil){
		glog.Fatal("%s",err)
		return
	}
	loginServer.DataServerXC = g_Entity.ConnectDataServer()
	g_Entity.Start("LoginServer",&loginServer)
	
}

func  Stop()  {
	g_Entity.Stop()
}