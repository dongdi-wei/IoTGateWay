package model

import (
	"IoTGateWay/base"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var(
	DBIot *gorm.DB			//全局使用的数据库连接句柄
	Logger *base.LogIot
)

func Init() {
	Logger = base.IotLogger
	DBIot, _ = gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",database_uname,database_passwd,database_ip,database_port,database_name))
	if DBIot != nil {
		Logger.Info("db connect success")
	}
}

func init()  {
	Init()
}