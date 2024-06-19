package auth

import (
	"backend/db"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(username string, password string) (error, int) {
	// Check if username already exists
	rows, err := db.Query("SELECT COUNT(username) as users FROM user WHERE username = ?", username)
	if err != nil {
		return err, 500
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return err, 500
		}
	}
	if count > 0 {
		return errors.New("username already exists"), 409
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err, 500
	}

	err = db.ExecStatement("INSERT INTO user (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
