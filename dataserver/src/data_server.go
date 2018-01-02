package dataserver

import (
	"flag"
	"../../server"
)

var (
	addr = flag.String("addr", "localhost:8810", "data server address")
)

type IDataServer struct {
	iserver.IServer
}

func Start(){
	flag.Parse()

	s := IDataServer{}
	s.Start("DataServer",new(DataServer),*addr)
}