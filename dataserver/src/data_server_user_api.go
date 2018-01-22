package dataserver

import (
	"context"
	"../../utils"
	"../../server"
	"../../common"
	"errors"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
)



//--------------RegisterUser------------//

type RegisterUserArgs struct {
	Account string
	Password string
}

type RegisterUserReply struct {
	iserver.IApiReply
	UserInfo 	User
}

func (self *DataServer) RegisterUser(ctx context.Context, args *RegisterUserArgs, reply *RegisterUserReply) error {
	glog.Infoln("RegisterUser user account=",args.Account)

	if (args.Account == "" || args.Password == ""){
		return errors.New("params is error "+args.Account+args.Password)
	}

	has, err := self.XormEngine.Where("account=?", args.Account).Get(&reply.UserInfo)

	if (has){
		reply.Msg = "账号名已有人使用"
		reply.Code = iserver.ApiCodeFail
		return nil
	}

	accout := args.Account
	password := utils.MD5([]byte(args.Password+common.Salt))

	User := new(User)
	User.Account = accout
	User.Password = password
	User.NickName = accout
	affected, err := self.XormEngine.Insert(User)
	if (affected != 1 || err != nil){
		reply.Msg = "注册失败"
		reply.Code = iserver.ApiCodeFail
	}else{
		reply.Code = iserver.ApiCodeSuccess
	}

    return err
}


//--------------QueryUser------------//

type QueryUserArgs struct {
	Account string
}

type QueryUserReply struct {
	iserver.IApiReply
	Has			bool
	UserInfo		User
}

func (self *DataServer) QueryUser(ctx context.Context, args *QueryUserArgs, reply *QueryUserReply) error {
	glog.Infoln("query user account=",args.Account)

	if (args.Account == ""){
		return errors.New("params is error "+args.Account)
	}

	has, err := self.XormEngine.Where("account=?", args.Account).Get(&reply.UserInfo)
	
	if (has){
		reply.Code = iserver.ApiCodeSuccess
	}else{
		reply.Code = iserver.ApiCodeFail
	}
	reply.Has = has
	return err
}