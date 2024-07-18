package db

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	ID              string `gorm:"primaryKey;size:36"`
	Email           string `gorm:"unique;not null"`
	IsEmailVerified bool   `gorm:"default:false"`
	HashedPassword  string `gorm:"not null"`
}

const (
	UserColumnID              ColumnName = "ID"
	UserColumnEmail           ColumnName = "Email"
	UserColumnIsEmailVerified ColumnName = "IsEmailVerified"
	UserColumnHashedPassword  ColumnName = "HashedPassword"
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.NewString()
	}
	return
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

func (h *dbHelper) UpdateUser(userID string, fields map[string]interface{}) error {
	result := h.db.Model(&User{}).Where("id = ?", userID).Updates(fields)
	return result.Error
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
