package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	Email           string `gorm:"unique;not null"`
	IsEmailVerified bool   `gorm:"default:false"`
	HashedPassword  string `gorm:"not null"`
}

func (h *dbHelper) InsertUsers(users []User) error {
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logrus.WithField("panic", r).Error("Transaction rolled back due to panic")
			tx.Rollback()
		}
	}()

	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (h *dbHelper) FindUserByEmail(email string) (*User, error) {
	var user User
	if err := h.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundError
		}
		return nil, err
	}
	return &user, nil
}
