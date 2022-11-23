package model

import "time"



type User struct {
	UserID   string
	Password string
}



type PersonalInfo struct {
	ID      int
	Name    string
	Phone   string
	Address string
}



type UserInfo struct {
	ID     int
	UserID string
}

type CustomerRegister struct {
	ID      int
	UserID   string
	Password string
	Name     string
	Phone    string
	Address  string
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