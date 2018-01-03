package dataserver

import (
	"flag"
	"../../server"
	"../../log"
)

var (
	addr = flag.String("addr", "localhost:8810", "data server address")
)

type IDataServer struct {
	iserver.IServer
}

func (self *IDataServer) Init(){
	log.Log(log.Debug,"IDataServerIDataServer")
}

func Start(){
	flag.Parse()

	s := IDataServer{}
	s.Init()
	s.Start("DataServer",new(DataServer),*addr)
}