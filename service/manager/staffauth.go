package manager

import (
	"errors"

	"mr.jackpot-backend/model"
)

type StafAuthService interface {
	AuthService

	CheckAuthority(user model.Staff) (int, error)
	CreateAccount(staff model.StaffRegister) error
	RemoveAccount(staffid int) error
}


func (m *StaffManager) CheckAuthority(user model.Staff) (userid int, err error) {
	userid, err = m.db.GetStaffID(user.Code)
	if err != nil {
		return
	}
	if userid == 0 {
		return 0, errors.New("NO ID")
	}
	return
}

func (m *StaffManager) CreateAccount(staff model.StaffRegister) error {
	return m.db.CreateAccount(staff)
}

func (m *StaffManager) RemoveAccount(staffid int) error {
	return m.db.SetAccounQuit(staffid)	
}
