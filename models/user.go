package models

import (
	"time"

	"github.com/wfanxin/go-service/db"
)

type User struct {
	ID        int
	Mobile    string    `gorm:"size:20"`
	Name      string    `gorm:"size:100"`
	Avatar    string    `gorm:"size:200"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

func GetUser(id int) (User, error) {
	var user User
	err := db.Mysql.First(&user, "id = ?", id).Error
	return user, err
}

func CreateUser() {
	var user User
	user.Mobile = "15880199707"
	user.Name = "xueg"
	db.Mysql.Create(&user)
}
