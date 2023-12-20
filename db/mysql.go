package db

import (
	"fmt"

	"github.com/wfanxin/go-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database)

	Mysql, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
