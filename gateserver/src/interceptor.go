package gateserver

import (
	"errors"
	"net/http"
	"../../server"
	"github.com/smallnest/rpcx/codec"
	"github.com/golang/glog"
)

func CheckToken(serverPath string,serviceMethod string,userId string,token string) bool{
	glog.Infoln("CheckToken",userId,token)
	return true
}


func Interceptor(r *http.Request) ([]byte,error){
	cc := &codec.MsgpackCodec{}
	args := iserver.IApiReply{}
	var err error

	for {
		serverPath := r.Header.Get(XServicePath) 
		serviceMethod := r.Header.Get(XServiceMethod)
		userId := r.Header.Get(XUserId)
		token := r.Header.Get(XToken)

		//不能接连数据库服务器
		if(serverPath == "DataServer"){
			args.Msg = "非法操作"
			args.Code = iserver.ApiCodeIllegalConnentServer
			err = errors.New("can non connect to the service DataServer")
			break
		}

		//检查Token
		if(!CheckToken(serverPath,serviceMethod,userId,token)){
			args.Msg = "token失效"
			args.Code = iserver.ApiCodeLoginInvalid
			err = errors.New("token is invalid")
			break
		}


		break
	}
	

	data, _ := cc.Encode(args)
	return data,err
}