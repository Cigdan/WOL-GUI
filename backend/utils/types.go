package utils

import "time"

type Device struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	MacAddress string     `json:"mac_address"`
	IpAddress  *string    `json:"ip_address"`
	LastOnline *time.Time `json:"last_online"`
	UserID     int        `json:"user_id"`
}

type UserData struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
