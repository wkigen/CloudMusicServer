package loginserver


import (
	"../../log"
	"../../server"
	"github.com/smallnest/rpcx/client"
)

type ServerEntity struct {
	iserver.IServer
}

type LoginServer struct{
	DataServerXC client.XClient
}

func Start(){
	loginServer := LoginServer{}
	entity := ServerEntity{}
	err := entity.Init()
	if(err != nil){
		log.Log(log.Fatel,"%s",err)
		return
	}
	loginServer.DataServerXC = entity.ConnectDataServer()
	entity.Start("LoginServer",&loginServer)
}