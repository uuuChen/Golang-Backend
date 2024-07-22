package db

import (
	"errors"
	"fmt"
	"glossika/internal/config"

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
	ListRecommendations() ([]Product, error)
}

type dbHelper struct {
	db *gorm.DB
}

func Init() (I, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config.AppConfig.MySQLUser,
		config.AppConfig.MySQLPassword,
		config.AppConfig.MySQLDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql db: %w", err)
	}

	err = db.AutoMigrate(&User{}, &Product{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	return &dbHelper{
		db: db,
	}, nil
}
