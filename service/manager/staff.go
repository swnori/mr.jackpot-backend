package manager

import "mr.jackpot-backend/database/db"



type StaffService interface {
	StafAuthService
	StaffInfoService
	AuthService
}

type StaffManager struct {
	db db.StaffLayer

	Client
	Member
}

var Staff = &StaffManager{}

func NewStaffManager() *StaffManager {
	Staff.Token = DefaultToken
	Staff.db = db.NewStaffDB()
	return Staff
}