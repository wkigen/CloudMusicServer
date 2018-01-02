package loginserver


import (
	"flag"
	"context"
	"errors"
	"../../log"
	"../../common"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8801", "login server address")
)

func auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == common.ServerToken {
		log.Log(log.Info,"gate server connect success")
		return nil
	}

	return errors.New("invalid token")
}

func Start(){
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("LoginServer", new(LoginServer), "")
	s.AuthFunc = auth
	s.Serve("tcp", *addr)
}