package dataserver

import (
	"time"
)

//--------User表数据结构------------//
type User struct {
    Id int64
	NickName string
	Account string
	Password string
    CreateTime time.Time `xorm:"created"`
}