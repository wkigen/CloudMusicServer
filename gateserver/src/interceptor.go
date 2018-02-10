package gateserver

import (
	"errors"
	"net/http"
	"../../server"
	"github.com/smallnest/rpcx/codec"
	"github.com/golang/glog"
)

type GateInterceptorReply struct{
	iserver.IApiReply
}



func CheckToken(userId string,token string) bool{
	glog.Infoln("CheckToken",userId,token)
	return true
}


func Interceptor(r *http.Request) ([]byte,error){
	cc := &codec.MsgpackCodec{}
	args := GateInterceptorReply{}
	var err error

	for {
		
		//不能接连数据库服务器
		if(r.Header.Get(XServicePath) == "DataServer"){
			args.Msg = "非法操作"
			args.Code = iserver.ApiCodeIllegalConnentServer
			err = errors.New("can non connect to the service DataServer")
			break
		}

		//检查Token
		if(!CheckToken(r.Header.Get(XUserId) ,r.Header.Get(XToken))){

			break
		}


		break
	}
	

	data, _ := cc.Encode(args)
	return data,err
}