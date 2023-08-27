package models

import (
	"Programa5/internal/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string    `gorm:"type:char(100);not null;uniqueIndex;size:100;"`
	Password  string    `gorm:"not null;type:char(64)"`
	LastLogin time.Time `gorm:""`
}

func CreateUser(email, password string) (*User, error) {
	if len(password) == 0 {
		return nil, errors.New("cannot create the user")
	}

	hashPassword := utils.HashSHA256(password)
	return &User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
		Email:     email,
		Password:  hashPassword,
		LastLogin: time.Time{},
	}, nil
}
