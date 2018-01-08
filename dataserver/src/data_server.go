package dataserver

import (
	"../../server"
	"fmt"
	"database/sql"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
)

type ServerEntity struct {
	iserver.IServer
}

type DataServer struct{
	DataBase *sql.DB
	Salt string
}

var g_Entity ServerEntity

func Init(s *DataServer) error{

	s.Salt = "cloudmusic"

	var err error
	databaseConf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",g_Entity.Config.DataBase.Accout,g_Entity.Config.DataBase.Password,
						g_Entity.Config.DataBase.Ip,g_Entity.Config.DataBase.Name)				
	s.DataBase, err = sql.Open(g_Entity.Config.DataBase.Type, databaseConf)
	
	return err
}



func Start(){
	dataServer := DataServer{}
	g_Entity = ServerEntity{}

	err := g_Entity.Init()
	if(err != nil){
		glog.Fatal(err)
		return
	}

	err = Init(&dataServer)
	if(err != nil){
		glog.Fatal(err)
		return
	}

	g_Entity.Start("DataServer",&dataServer)
}

func Stop(){
	g_Entity.Stop()
}