package loginserver


import (
	"../../log"
	"../../server"
)

type ServerEntity struct {
	iserver.IServer
}

type LoginServer struct{

}

func Start(){
	loginServer := LoginServer{}
	entity := ServerEntity{}
	err := entity.Init()
	if(err != nil){
		log.Log(log.Fatel,"%s",err)
		return
	}
	entity.Start("LoginServer",&loginServer)
}