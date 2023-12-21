package db

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func init() {
	file, err := ini.Load("./config/env.ini")
	if err != nil {
		fmt.Printf("env.ini读取失败：%v\n", err)
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		file.Section("mysql").Key("username"),
		file.Section("mysql").Key("password"),
		file.Section("mysql").Key("host"),
		file.Section("mysql").Key("port"),
		file.Section("mysql").Key("database"))

	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql连接失败：%v", err))
	}
}
