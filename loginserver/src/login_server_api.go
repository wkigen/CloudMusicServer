package loginserver

import (
	"context"
	"../../log"
	"../../server"
	"../../dataserver/src"
)

//--------------RegisterUser------------//
type RegisterUserArgs struct {
	Accout string
	Password string
}

type RegisterUserReply struct {
	iserver.IApiReply
}

func (self *LoginServer) RegisterUser(ctx context.Context, args *RegisterUserArgs, reply *RegisterUserReply) error {

	qArgs := &dataserver.RegisterUserArgs{}
	qReply := &dataserver.QueryUserReply{}

	qArgs.Accout = args.Accout
	qArgs.Password = args.Password

	err := self.DataServerXC.Call(ctx,"RegisterUser",qArgs,qReply)

	reply.Code = qReply.Code
	reply.Msg = qReply.Msg

    return err
}


//--------------QueryUser------------//

type LoginArgs struct{
	Accout string
	Passwrod string
}

type LoginReply struct{
	iserver.IApiReply
	Id			int32
	Accout 		string
	NickName 	string
	Token 		string
}

func (self *LoginServer) Login(ctx context.Context, args *LoginArgs, reply *LoginReply) error {
	log.Log(log.Debug,"--------login--------")
	log.Log(log.Debug,"user accout:"+args.Accout)
	log.Log(log.Debug,"user password:"+args.Passwrod)

	qArgs := &dataserver.QueryUserArgs{}
	qReply := &dataserver.QueryUserReply{}

	qArgs.Accout = args.Accout

	err := self.DataServerXC.Call(ctx,"QueryUser",qArgs,qReply)

	if(err == nil){
		if(qReply.Id != -1){
			reply.Code = 0
			reply.Id = qReply.Id
			reply.Accout = qReply.Accout
			reply.NickName = qReply.NickName
			reply.Token = "100"
		}else{
			reply.Code = 1
			reply.Msg = qReply.Msg
		}
	}

    return err
}


