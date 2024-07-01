package auth

import (
	"backend/utils"
	"backend/utils/logger"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(driver *sql.DB, username string, password string) (error, int) {
	// Check if username already exists
	row, err := utils.QueryOne(driver, "SELECT COUNT(username) as users FROM user WHERE username = ?", username)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return err, 500
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		logger.Error("Error scanning row: " + err.Error())
		return err, 0
	}

	if count > 0 {
		logger.Warning("Username: " + username + " already exists")
		return errors.New("username already exists"), 409
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Error("Error hashing password: " + err.Error())
		return err, 500
	}

	_, err = utils.ExecStatement(driver, "INSERT INTO user (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		logger.Error("Error inserting user: " + err.Error())
		return err, 500
	}
	return nil, 200
}

func GenerateToken(id int, username string) (string, error) {
	key, isKey := os.LookupEnv("JWT_KEY")
	if !isKey {
		err := generateKey()
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
		ID:       id,
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

func generateKey() error {
	log.Println("JWT_KEY environment variable not set, new key will be generated")
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		return err
	}
	err = os.Setenv("JWT_KEY", base64.URLEncoding.EncodeToString(secret))
	if err != nil {
		return err
	}
	return nil
}

func Login(driver *sql.DB, username string, password string) (string, error, int) {
	var dbId int
	var dbUsername string
	var dbPassword string

	row, err := utils.QueryOne(driver, "SELECT id, username, password FROM user WHERE username = ?", username)
	if err != nil {
		return "", err, 400
	}
	err = row.Scan(&dbId, &dbUsername, &dbPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warning("User " + username + " not found")
			return "", errors.New("invalid Credentials"), 401
		}
		logger.Error("Error scanning row: " + err.Error())
		return "", err, 500
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		logger.Warning("Invalid password for user " + username)
		return "", errors.New("invalid Credentials"), 401
	}

	token, err := GenerateToken(dbId, username)
	if err != nil {
		logger.Error("Error generating token: " + err.Error())
		return "", err, 500
	}
	return token, nil, 200
}

func ValidateToken(reqToken string) (*Claims, error, int) {
	key, isKey := os.LookupEnv("JWT_KEY")
	if !isKey {
		err := generateKey()
		if err != nil {
			return nil, err, 500
		}
		key = os.Getenv("JWT_KEY")
	}

	decodedKey, err := base64.URLEncoding.DecodeString(key)
	if err != nil {
		logger.Error("Error decoding key: " + err.Error())
		return nil, err, 500
	}

	token, err := jwt.ParseWithClaims(reqToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return decodedKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			logger.Warning("Invalid token signature")
			return nil, errors.New("invalid token signature"), 401
		}
		logger.Error("Error parsing token: " + err.Error())
		return nil, err, 500
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Check if the token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			logger.Warning("Token is expired: " + claims.ExpiresAt.String())
			return nil, errors.New("token is expired"), 401
		}
		return claims, nil, 200
	}

	logger.Warning("Invalid token supplied: " + reqToken)
	return nil, errors.New("invalid token supplied"), 401
}
