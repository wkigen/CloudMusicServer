package dataserver

import (
	"../../server"
	"../../log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ServerEntity struct {
	iserver.IServer
}

type DataServer struct{
	DataBase *sql.DB
	Salt string
}

func Init(s *DataServer,entity ServerEntity) error{

	s.Salt = "cloudmusic"

	var err error
	databaseConf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",entity.Config.DataBase.Accout,entity.Config.DataBase.Password,
						entity.Config.DataBase.Ip,entity.Config.DataBase.Name)				
	s.DataBase, err = sql.Open(entity.Config.DataBase.Type, databaseConf)
	
	return err
}

func Start(){
	dataServer := DataServer{}
	entity := ServerEntity{}

	err := entity.Init()
	if(err != nil){
		log.Log(log.Fatel,"%s",err)
		return
	}

	err = Init(&dataServer,entity)
	if(err != nil){
		log.Log(log.Fatel,"%s",err)
		return
	}

	entity.Start("DataServer",&dataServer)
}