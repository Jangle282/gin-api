package models

import (
	"errors"
	"gin-api/database"
	"gin-api/internal/hashing"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := hashing.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := hashing.CheckPassword(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("Credentials invalid.")
	}

	return nil
}

func GetUserById(id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id=?`
	//When you know you will get back only 1 row.
	row := db.DB.QueryRow(query, id)
	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		// type return is *Event = null pointer for Event is nil,
		return nil, err
	}

	return &user, nil
}
