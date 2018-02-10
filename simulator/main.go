package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	Id			int64
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

type ChangNickNameArgs struct {
	NewNickName string
}

type ChangNickNameReply struct {
	Code int
	Msg string
}

func Register(index int){
	cc := &codec.MsgpackCodec{}

	args := &RegisterUserArgs{
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

	var mId = 20000+index
	h := req.Header
	h.Set(gateway.XMessageID, strconv.Itoa(mId))
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

	log.Printf("%s , %s ,%d,%s", args.Account, args.Password, reply.Code,reply.Msg)
}

var userId int64
var token string
func Login(index int){

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

	var mId = 10000+index
	h := req.Header
	h.Set(gateway.XMessageID,strconv.Itoa(mId))
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

	userId = reply.Id
	token = reply.Token

	log.Printf(" login %s , %s ,%s,%s", args.Account, args.Password, reply.Token,reply.Msg)

	ChangNickName(reply.Id,reply.Token)
}

func ChangNickName(userId int64,token string){
	cc := &codec.MsgpackCodec{}

	args := &ChangNickNameArgs{
		NewNickName:"fff",
	}

	data, _ := cc.Encode(args)
	b := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8701/",b )
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	var mId = 11000
	h := req.Header
	h.Set(gateway.XMessageID,strconv.Itoa(mId))
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "LoginServer")
	h.Set(gateway.XServiceMethod, "ChangNickName")
	h.Set("UserId", strconv.FormatInt(userId,10))
	h.Set("Token", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &ChangNickNameReply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf(" ChangNickName %d , %s ", reply.Code,reply.Msg)

}

func TestDataServer(index int){
	cc := &codec.MsgpackCodec{}

	args := &LoginArgs{
		Account: "1",
		Password:"1",
	}

	data, _ := cc.Encode(args)
	b := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8701/",b )
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	var mId = 10000+index
	h := req.Header
	h.Set(gateway.XMessageID,strconv.Itoa(mId))
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "DataServer")
	h.Set(gateway.XServiceMethod, "Login")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &LoginReply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf(" TestDataServer %d , %s ", reply.Code,reply.Msg)

}

func main() {
	
	ch := make(chan int) 
	for index := 0; index < 1; index++ {
		log.Printf("%d",index)
		go Login(index)
		//go Register(index)
		//go TestDataServer(index)
	}
	<-ch
}
