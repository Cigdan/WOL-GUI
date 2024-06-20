package auth

import (
	"backend/db"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(username string, password string) (error, int) {
	// Check if username already exists
	row, err := db.QueryOne("SELECT COUNT(username) as users FROM user WHERE username = ?", username)
	if err != nil {
		return err, 500
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return err, 0
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

func GenerateToken(id int, username string) (string, error) {
	key, isKey := os.LookupEnv("JWT_KEY")
	if !isKey {
		err := InitAuth()
		if err != nil {
			return "", err
		}
		key = os.Getenv("JWT_KEY")
	}

	decodedKey, err := base64.URLEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(decodedKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func InitAuth() error {
	_, keyExists := os.LookupEnv("JWT_KEY")
	if !keyExists {
		log.Println("JWT_KEY environment variable not set, new key will be generated")
		secret := make([]byte, 32)
		_, err := rand.Read(secret)
		if err != nil {
			log.Println("Error generating random secret", err)
			return err
		}
		err = os.Setenv("JWT_KEY", base64.URLEncoding.EncodeToString(secret))
		if err != nil {
			return err
		}
	}
	return nil
}
func Login(username string, password string) (string, error, int) {
	var dbId int
	var dbUsername string
	var dbPassword string

	row, err := db.QueryOne("SELECT id, username, password FROM user WHERE username = ?", username)
	if err != nil {
		return "", err, 400
	}
	err = row.Scan(&dbId, &dbUsername, &dbPassword)
	if err != nil {
		return "", err, 500
	}
	if dbUsername != username {
		return "", errors.New("username doesn't exist"), 400
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return "", errors.New("wrong password"), 401
	}

	token, err := GenerateToken(dbId, username)
	if err != nil {
		return "", err, 500
	}
	return token, nil, 200
}

func ValidateToken(reqToken string) (*Claims, error, int) {
	token, err := jwt.ParseWithClaims(reqToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token supplied"), 401
		}
		return nil, err, 500
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil, 200
	}
	return nil, errors.New("invalid token supplied"), 401

}
