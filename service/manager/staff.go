package manager

import "mr.jackpot-backend/database/db"



type StaffService interface {
	StafAuthService
	StaffInfoService
}

type StaffManager struct {
	db db.StaffLayer
	Manager
}

var Staff = &StaffManager{
	db: db.NewStaffDB(),
}

func NewStaffManager() *StaffManager {
	return Staff
}