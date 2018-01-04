package dataserver

import (
	"../../log"
	"context"
	"github.com/go-sql-driver/mysql"
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
	password := args.Password
	log.Log(log.Debug,RegisterUser)
	if(accout != "" && password != ""){
		stmt, err :=self.DataBase.Prepare("INSERT INTO user VALUES (?,?,?,?)")
		defer stmt.Close()
		if err == nil {
			stmt.Exec(0,accout,accout,password)
			reply.Status = 0
			return nil
		}
	}
	reply.Status = 1
    return nil
}


//--------------QueryUser------------//

type QueryUserArgs struct {

}

type QueryUserReply struct {

}

func (self *DataServer) QueryUser(ctx context.Context, args *QueryUserArgs, reply *QueryUserReply) error {

    return nil
}