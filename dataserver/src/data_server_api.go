package dataserver

import (
	"../../log"
	"../../utils"
	"errors"
	"context"
	_ "github.com/go-sql-driver/mysql"
)

//--------------RegisterUser------------//

type RegisterUserArgs struct {
	Accout string
	Password string
}

type RegisterUserReply struct {
	Status int
}

func (self *DataServer) RegisterUser(ctx context.Context, args *RegisterUserArgs, reply *RegisterUserReply) error {
	accout := args.Accout

	qArgs := QueryUserArgs{}
	qArgs.Accout = accout
	qReply := QueryUserReply{}
	err := self.QueryUser(ctx,&qArgs,&qReply)
	if(err == nil ||err.Error() == "sql: no rows in result set"){
		if( qReply.Id == -1){
			password := utils.MD5([]byte(args.Password+self.Salt))
			if(accout != "" && password != ""){
				stmt, err := self.DataBase.Prepare("INSERT INTO user VALUES (?,?,?,?)")
				defer stmt.Close()
				if( err == nil ){
					res,err := stmt.Exec(0,accout,accout,password)
					if(err == nil){
						reply.Status = 0
						log.Log(log.Debug,"user registration success,(%s) %d",accout,res.LastInsertId)
						return nil
					}
				}
			}
		}else{
			err = nil
			log.Log(log.Debug,"user registration fail，error :has same accout in db")
		}
	}
	log.Log(log.Debug,"user registration fail， error :%s",err)
	reply.Status = 1
    return err
}


//--------------QueryUser------------//

type QueryUserArgs struct {
	Accout string
}

type QueryUserReply struct {
	Id			int32
	Accout 		string
	NickName 	string
	Password	string
}

func (self *DataServer) QueryUser(ctx context.Context, args *QueryUserArgs, reply *QueryUserReply) error {
	reply.Id = -1
	rows := self.DataBase.QueryRow("SELECT * FROM user WHERE account = ?",args.Accout)
	err := rows.Scan(&reply.Id, &reply.Accout, &reply.NickName, &reply.Password)
	log.Log(log.Debug,"%d %s %s %s",reply.Id, reply.Accout, reply.NickName, reply.Password)
	if(err == nil){
		return nil
	}
	log.Log(log.Warn,"query user (%s) is error,%s",args.Accout,err)
    return err
}