package loginserver

import (
	"context"
	"../../log"
)

type LoginServer int

type LoginArgs struct{
	Name string
	Passwrod string
}

type LoginReply struct{
	Token string
}

func (self *LoginServer) Login(ctx context.Context, args *LoginArgs, reply *LoginReply) error {
	log.Log(log.Debug,"--------login--------")
	log.Log(log.Debug,"user name:"+args.Name)
	log.Log(log.Debug,"user password:"+args.Passwrod)
    reply.Token = "1"
    return nil
}