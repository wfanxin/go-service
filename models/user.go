package models

import (
	"database/sql"
	"time"

	"github.com/wfanxin/go-service/db"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string {
	return "users"
}

func GetUser(id int) (User, error) {
	var user User
	err := db.Mysql.First(&user, "id = ?", id).Error
	return user, err
}
