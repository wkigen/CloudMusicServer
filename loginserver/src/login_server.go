package loginserver


import (
	"../../server"
)

type ILoginServer struct {
	iserver.IServer
}

func Start(){
	s := ILoginServer{}
	s.Start("LoginServer",new(LoginServer))
}