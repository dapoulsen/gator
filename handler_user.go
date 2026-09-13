package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dapoulsen/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 || len(cmd.Args) > 1 {
		return errors.New("A username is needed")
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("user isn't in the database")
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("user have been switched to: ", cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("No user was provided")
	}

	name := cmd.Args[0]

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	dbUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("User was successfully registered")
	fmt.Println(dbUser)

	return nil

}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error getting users")
		return err
	}

	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Println("*  " + user.Name + " (current)")
		} else {
			fmt.Println("* ", user.Name)
		}
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("Error resetting users table", err)
		os.Exit(1)
	}
	fmt.Println("users table was successfully reset")
	return nil
}
