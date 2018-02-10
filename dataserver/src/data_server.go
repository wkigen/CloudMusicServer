package dataserver

import (
	"../../server"
	"fmt"
	"time"
	"sync"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type DataServer struct{
	Base 		iserver.IServer
	XormEngine 	*xorm.Engine
	SyMu       	sync.RWMutex
}

var g_DataServer DataServer

func PingDB(){
	for {
		if(g_DataServer.XormEngine != nil){
			err := g_DataServer.XormEngine.Ping()
			if (err != nil){
				glog.Errorln("can not ping the db",err)
			}
		}
		time.Sleep(time.Minute * 10)
	}
}

func Init() error{
	var err error
	databaseConf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",g_DataServer.Base.Config.DataBase.Accout,
						g_DataServer.Base.Config.DataBase.Password,
						g_DataServer.Base.Config.DataBase.Ip,
						g_DataServer.Base.Config.DataBase.Name)				
	g_DataServer.XormEngine, err = xorm.NewEngine(g_DataServer.Base.Config.DataBase.Type,databaseConf)

	go PingDB()
	return err
}

func Start(){
	g_DataServer = DataServer{}

	err := g_DataServer.Base.Init()
	if(err != nil){
		glog.Fatal(err)
		return
	}

	err = Init()
	if(err != nil){
		glog.Fatal(err)
		return
	}

	g_DataServer.Base.Start("DataServer",&g_DataServer)
}

func Stop(){
	g_DataServer.Base.Stop()
	if(g_DataServer.XormEngine != nil){
		g_DataServer.XormEngine.Close()
	}
}