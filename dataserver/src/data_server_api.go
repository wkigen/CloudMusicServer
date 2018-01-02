package dataserver

import (
	"context"
)

type DataServer int

type QueryUserArgs struct {

}

type QueryUserReply struct {

}

func (self *DataServer) QueryUser(ctx context.Context, args *QueryUserArgs, reply *QueryUserReply) error {

    return nil
}