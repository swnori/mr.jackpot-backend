package manager

import (
	"errors"

	"mr.jackpot-backend/model"
)

type StafAuthService interface {
	CheckAuthority(user model.Staff) error
	CreateAccount(staff model.StaffInfo) error
	RemoveAccount(staffid int) error
}


func (m *StaffManager) CheckAuthority(user model.Staff) error {
	userid, err := m.db.GetStaffID(user.Code)
	if err != nil {
		return err
	}
	if userid == 0 {
		return errors.New("NO ID")
	}
	return nil
}

func (m *StaffManager) CreateAccount(staff model.StaffInfo) error {
	return m.db.CreateAccount(staff)
}

func (m *StaffManager) RemoveAccount(staffid int) error {
	return m.db.SetAccounQuit(staffid)	
}
