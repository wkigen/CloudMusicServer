package dataserver

import (
	"../../server"
)

type IDataServer struct {
	iserver.IServer
}

func (self *IDataServer) Init(){
	
}

func Start(){
	s := IDataServer{}
	s.Init()
	s.Start("DataServer",new(DataServer))
}