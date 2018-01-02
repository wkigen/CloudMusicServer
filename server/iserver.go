package iserver

import (
	"context"
	"errors"
	"../log"
	"../common"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/protocol"
)

type IServer struct{
	Name string
}

func (self *IServer)auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == common.ServerToken {
		log.Log(log.Info,"Authentication Success")
		return nil
	}
	log.Log(log.Info,"Authentication Fail")
	return errors.New("invalid token")
}

func (self *IServer) Start(name string,rcvr interface{},addr string ) {
	self.Name = name

	s := server.NewServer()
	s.RegisterName(name, rcvr, "")
	s.AuthFunc = self.auth
	s.Serve("tcp", addr)
} 