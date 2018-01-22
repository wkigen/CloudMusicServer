package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rpcx-ecosystem/rpcx-gateway"

	"github.com/smallnest/rpcx/codec"
)

type LoginArgs struct{
	Account string
	Password string
}

type LoginReply struct{
	Code int
	Msg string
	Id			int32
	Accout 		string
	NickName 	string
	Token 		string
}

type RegisterUserArgs struct {
	Account string
	Password string
}

type RegisterUserReply struct {
	Code int
	Msg string
}

func Register(){
	cc := &codec.MsgpackCodec{}

	args := &RegisterUserArgs{
		Account: "woshishui004",
		Password:"123456",
	}

	data, _ := cc.Encode(args)
	b := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8701/",b )
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "LoginServer")
	h.Set(gateway.XServiceMethod, "RegisterUser")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	// handle http response
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &RegisterUserReply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf("%s , %s ,%s", args.Account, args.Password, reply.Code,reply.Msg)
}

func Login(){
	
	cc := &codec.MsgpackCodec{}

	args := &LoginArgs{
		Account: "woshishui001",
		Password:"123456",
	}

	data, _ := cc.Encode(args)
	b := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8701/",b )
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "LoginServer")
	h.Set(gateway.XServiceMethod, "Login")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	// handle http response
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &LoginReply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf("%s , %s ,%s,%s", args.Account, args.Password, reply.Token,reply.Msg)
}

func main() {
	
	ch := make(chan int) 

	for index := 0; index <= 1000; index++ {
		log.Printf("%d",index)
		go 	Login()
	}

	<-ch
	// Register()
}
