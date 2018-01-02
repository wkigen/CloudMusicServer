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
	Name string
	Passwrod string
}

type LoginReply struct{
	Token string
}

func main() {
	cc := &codec.MsgpackCodec{}

	args := &LoginArgs{
		Name: "10",
		Passwrod:"20",
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

	log.Println(h.Get("X-RPCX-ServicePath"))

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

	log.Printf("%s , %s ,%s", args.Name, args.Passwrod, reply.Token)
}
