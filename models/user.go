package models

import (
	"api/db"
	"api/utils"
	"errors"
	// "fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, utils.HashPassword(user.Password))
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user *User) ValidateCreds() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials")
	}
	// fmt.Printf("-----retrieved pass: %s\n", retrievedPassword)

	// compare user password with hashed password
	valid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	// fmt.Println("-----valid", valid)
	if valid {
		return nil
	}
	return errors.New("Invalid credentials")
}
