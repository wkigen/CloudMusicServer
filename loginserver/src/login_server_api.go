package loginserver

import (
	"context"
	"github.com/golang/glog"
	"../../server"
	"../../dataserver/src"
	"../../common"
	"../../utils"
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
	Password string
}

type LoginReply struct{
	iserver.IApiReply
	Id			int32
	Accout 		string
	NickName 	string
	Token 		string
}

func (self *LoginServer) Login(ctx context.Context, args *LoginArgs, reply *LoginReply) error {

	qArgs := &dataserver.QueryUserArgs{}
	qReply := &dataserver.QueryUserReply{}

	qArgs.Accout = args.Accout

	err := self.DataServerXC.Call(ctx,"QueryUser",qArgs,qReply)

	if(err == nil){
		if(qReply.Id != -1){
			psw := utils.MD5([]byte(args.Password+common.Salt)) 
			if(psw == qReply.Password){
				glog.Infoln("--------login--------")
				glog.Infoln("user accout:"+args.Accout)
				glog.Infoln("user password:"+args.Password)

				reply.Code = iserver.ApiCodeSuccess
				reply.Id = qReply.Id
				reply.Accout = qReply.Accout
				reply.NickName = qReply.NickName
				reply.Token = "100"
			}else{
				reply.Code = iserver.ApiCodeFail
				reply.Msg = "密码错误"
			}
		}else{
			reply.Code = iserver.ApiCodeFail
			reply.Msg = qReply.Msg
		}
	}

    return err
}


