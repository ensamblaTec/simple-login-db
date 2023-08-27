package controller

import (
	"Programa5/db/orm"
	"Programa5/internal/models"
	"Programa5/internal/utils"
	"errors"
	"log"
)

func RegisterUser(user *models.User) error {
	if err := orm.RegisterUser(user); err != nil {
		log.Println("[CONTROLLER FAILED]: Register user")
		return err
	}
	log.Println("[CONTROLLER SUCCESS]: Register user")
	return nil
}

func Login(user *models.User) error {
	newUser, err := orm.GetUser(user)
	if err != nil {
		log.Println("[CONTROLLER FAILED]: Login user")
		return err
	}
	if user.Email != newUser.Email && newUser.Password != utils.HashSHA256(user.Password) {
		log.Println("[CONTROLLER FAILED]: Login user")
		return errors.New("invalid credentials")
	}
	log.Println("[CONTROLLER SUCCESSFUL]: Login user")
	return nil
}
