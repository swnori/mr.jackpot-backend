package authority

import "mr.jackpot-backend/service/manager"




type StaffAuthService interface {
	SigningService
}

type StaffAuthHandler struct {
	m manager.StaffService
}

func NewStaffAuthHandler() *StaffAuthHandler {
	return &StaffAuthHandler{
		m: manager.NewStaffManager(),
	}
}