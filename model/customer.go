package model

import "time"



type User struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}



type PersonalInfo struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}



type UserInfo struct {
	ID     int
	UserID string
}

type CustomerRegister struct {
	ID      int
	UserID   string `json:"userID"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type CustomerInfo struct {
	ID      int
	UserID  string
	Name    string
	Phone   string
	Address string

	CreatedAt  time.Time
	Rating     int
	OrderCount int
	Paid       int
}