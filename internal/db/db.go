package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var NotFoundError = errors.New("record not found")

type ColumnName string

func (name ColumnName) ToString() string {
	return string(name)
}

type I interface {
	InsertUsers(users []User) error
	UpdateUser(userID string, fields map[string]interface{}) error
	FindUserByEmail(email string) (*User, error)
}

type dbHelper struct {
	db *gorm.DB
}

func Init() (I, error) {
	dsn := "root:rootpassword@tcp(db:3306)/mydb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql db: %w", err)
	}

	db.AutoMigrate(&User{})

	return &dbHelper{
		db: db,
	}, nil
}
