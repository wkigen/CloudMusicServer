package loginserver


import (
	"flag"
	"../../server"
)

var (
	addr = flag.String("addr", "localhost:8801", "login server address")
)

type ILoginServer struct {
	iserver.IServer
}

func Start(){
	flag.Parse()

	s := ILoginServer{}
	s.Start("LoginServer",new(LoginServer),*addr)
}