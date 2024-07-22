package users

import (
	"fmt"
	"glossika/internal/db"
)

type User struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

func (u *User) ToDB() (*db.User, error) {
	hashedPassword, err := HashPlainPassword(u.PlainPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &db.User{
		Email:          u.Email,
		HashedPassword: hashedPassword,
	}, nil
}
