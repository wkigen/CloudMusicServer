package loginserver


import (
	"../../server"
	"github.com/golang/glog"
)

type LoginServer struct{
	Base 			iserver.IServer
}

var g_LoginServer LoginServer

func Start(){
	glog.Infoln("LoginServer is start")
	g_LoginServer= LoginServer{}
	err := g_LoginServer.Base.Init()
	if(err != nil){
		glog.Fatal("%s",err)
		return
	}
	g_LoginServer.Base.ConnectDataServer()
	g_LoginServer.Base.ConnectRedis()
	g_LoginServer.Base.Start("LoginServer",&g_LoginServer)
}

func  Stop()  {
	g_LoginServer.Base.Stop()
}