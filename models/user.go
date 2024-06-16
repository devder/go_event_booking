package models

import (
	"errors"
	"fmt"

	"github.com/devder/go_event_booking/db"
	"github.com/devder/go_event_booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("err hashing password: ", err)
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var foundPassword string
	err := row.Scan(&foundPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPassword(foundPassword, u.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
