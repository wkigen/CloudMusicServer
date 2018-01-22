package dataserver

import (
	"../../server"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type ServerEntity struct {
	iserver.IServer
}

type DataServer struct{
	XormEngine *xorm.Engine
}

var g_Entity ServerEntity
var g_DataServer DataServer

func Init() error{
	var err error
	databaseConf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",g_Entity.Config.DataBase.Accout,g_Entity.Config.DataBase.Password,
						g_Entity.Config.DataBase.Ip,g_Entity.Config.DataBase.Name)				
	g_DataServer.XormEngine, err = xorm.NewEngine(g_Entity.Config.DataBase.Type,databaseConf)
	return err
}

func Start(){
	g_DataServer = DataServer{}
	g_Entity = ServerEntity{}

	err := g_Entity.Init()
	if(err != nil){
		glog.Fatal(err)
		return
	}

	err = Init()
	if(err != nil){
		glog.Fatal(err)
		return
	}

	g_Entity.Start("DataServer",&g_DataServer)
}

func Stop(){
	g_Entity.Stop()
	if(g_DataServer.XormEngine != nil){
		g_DataServer.XormEngine.Close()
	}
}