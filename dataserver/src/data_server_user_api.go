package dataserver

import (
	"../../utils"
	"context"
	"../../server"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
)

//--------------RegisterUser------------//

type RegisterUserArgs struct {
	Accout string
	Password string
}

type RegisterUserReply struct {
	iserver.IApiReply

}

func (self *DataServer) RegisterUser(ctx context.Context, args *RegisterUserArgs, reply *RegisterUserReply) error {
	accout := args.Accout

	qArgs := QueryUserArgs{}
	qArgs.Accout = accout
	qReply := QueryUserReply{}
	err := self.QueryUser(ctx,&qArgs,&qReply)
	if(err == nil){
		if( qReply.Id == -1){
			password := utils.MD5([]byte(args.Password+self.Salt))
			if(accout != "" && password != ""){
				stmt, err := self.DataBase.Prepare("INSERT INTO user VALUES (?,?,?,?)")
				defer stmt.Close()
				if( err == nil ){
					res,err := stmt.Exec(0,accout,accout,password)
					if(err == nil){
						reply.Code = 0
						reply.Msg = "注册成功"
						glog.Infoln("user registration success,(%s) %d",accout,res.LastInsertId)
						return nil
					}
				}
			}
		}else{
			err = nil
			reply.Msg = "账号名已有人使用"
		}
	}else{
		reply.Msg = "注册失败"
	}
	glog.Infoln("user registration fail， error :%s",err)
	reply.Code = 1
    return err
}


//--------------QueryUser------------//

type QueryUserArgs struct {
	Accout string
}

type QueryUserReply struct {
	iserver.IApiReply
	Id			int32
	Accout 		string
	NickName 	string
	Password	string
}

func (self *DataServer) QueryUser(ctx context.Context, args *QueryUserArgs, reply *QueryUserReply) error {
	rows := self.DataBase.QueryRow("SELECT * FROM user WHERE account = ?",args.Accout)
	err := rows.Scan(&reply.Id, &reply.Accout, &reply.NickName, &reply.Password)
	
	if(err == nil){
		glog.Infoln("QueryUser %d %s %s %s",reply.Id, reply.Accout, reply.NickName, reply.Password)
		reply.Code = 0
	}else{
		reply.Code = 1
		if(err.Error() == "sql: no rows in result set"){
			reply.Id = -1
			reply.Msg = "找不到该用户"
			err = nil
		}
		glog.Infoln("query user (%s) is error,%s",args.Accout,err)
	}
    return err
}