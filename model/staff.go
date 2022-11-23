package model

import "time"




 
type Staff struct {
	Code    string
	Name    string
	Role    string
	Part    string
}

type StaffInfo struct {
	ID    int
	Code  string
	Name  string
	Role  string
	Part  string
	Score int
	CreatedAt time.Time
}

type StaffRegister struct {
	Code string
	Name string
	RoleID int
}