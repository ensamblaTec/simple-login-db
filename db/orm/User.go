package orm

import (
	"Programa5/db"
	"Programa5/internal/models"
	"log"
)

// RegisterUser CreateUser Create a user
func RegisterUser(user *models.User) error {
	// init connection
	db.Connection()
	if err := db.DB.Select("email", "password", "created_at").Create(&user).Error; err != nil {
		log.Println("[ORM FAILED]: Register user")
		return err
	}
	log.Println("[ORM SUCCESS]: Register user")
	return nil
}

// GetUser Login with account
func GetUser(user *models.User) (*models.User, error) {
	db.Connection()
	if err := db.DB.Where("email", user.Email).Find(&user).Error; err != nil {
		log.Println("[ORM FAILED]: Login user")
		return nil, err
	}
	log.Println("[ORM SUCCESS]: Login user")
	return user, nil
}
