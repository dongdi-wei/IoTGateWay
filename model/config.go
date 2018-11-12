package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var(
	DBIot *gorm.DB			//全局使用的数据库连接句柄
)

func Init() {
	DBIot, err := gorm.Open("mysql",fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",database_uname,database_passwd,database_ip,database_port,database_name))
	if err != nil {
		fmt.Println("connect db error,%v", err)
	}
	defer DBIot.Close()
	if DBIot != nil {
		fmt.Println("db connect success")
	}
}
