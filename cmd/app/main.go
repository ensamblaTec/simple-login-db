package main

import (
	"Programa5/db"
	"Programa5/db/controller"
	"Programa5/internal/models"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

/*
email:	example@example.com
pass:	awacateConAwa

email:	admin@example.com
pass:	1234
*/
func main() {
	app := &cli.App{
		Name:        "simple login/register",
		Usage:       "register or login an user",
		Description: "This program keep an user into a database. Also you can login with any user.",
		Authors: []*cli.Author{{
			Name: "@ensamblaTec",
		}},
		Commands: []*cli.Command{
			{
				Name:        "register",
				Description: "Register an user",
				Aliases:     []string{"r"},
				Usage:       "type an email and password",
				Action: func(context *cli.Context) error {
					// check valid context
					if err := validInputs(context); err != nil {
						return err
					}
					email, password := context.Args().Get(0), context.Args().Get(1)
					// create user model
					user, err := models.CreateUser(email, password)
					if err != nil {
						return err
					}
					// create user on controller
					err = controller.RegisterUser(user)
					if err != nil {
						return err
					}
					// print user created
					fmt.Printf("====== User Created ======\nEmail: %s\nPassword: %s\nCreatedAt: %s\n", user.Email, user.Password, user.CreatedAt.Format("2006/01/02 15:04:05"))
					return nil
				},
			},
			{
				Name:        "login",
				Description: "Login with a email and password",
				Aliases:     []string{"l"},
				Usage:       "Type email and password",
				Action: func(context *cli.Context) error {
					// check valid context
					if err := validInputs(context); err != nil {
						return err
					}
					email, password := context.Args().Get(0), context.Args().Get(1)
					// create user model
					user, err := models.CreateUser(email, password)
					if err != nil {
						return err
					}
					// create user on controller
					err = controller.Login(user)
					if err != nil {
						return err
					}
					// print user created
					fmt.Printf("====== Login Successful ======\nEmail: %s", user.Email)
					return nil
				},
			},
			{
				Name:        "migrate",
				Description: "Migrate changes in code to database",
				Aliases:     []string{"m"},
				Usage:       "Not necessary args",
				Action: func(context *cli.Context) error {
					db.Connection()
					fmt.Println("Migrating...")
					time.Sleep(1 * time.Second)
					err := db.DB.AutoMigrate(&models.User{})
					if err != nil {
						log.Panic(err)
					}
					fmt.Println("Successful migrate")
					return nil
				},
			},
		},
		Action: func(*cli.Context) error {
			fmt.Println("Need arguments...")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func validInputs(context *cli.Context) error {
	if context.NArg() == 0 {
		return errors.New("missing arguments")
	}
	if context.NArg() > 2 {
		return errors.New("expected 2 arguments")
	}
	return nil
}
