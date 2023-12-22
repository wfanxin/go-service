package dao

import (
	"github.com/wfanxin/go-service/db"
	"github.com/wfanxin/go-service/models"
)

func Migrate() {
	db.Mysql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}
