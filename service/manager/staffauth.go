package manager

import (
	"errors"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type StafAuthService interface {
	AuthService

	CheckAuthority(user model.Staff) (int, error)
	CreateAccount(staff model.StaffRegister) (string, error)
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

func (m *StaffManager) CreateAccount(staff model.StaffRegister) (string, error) {

	var prefix string
	if staff.RoleID == 1 {
		prefix = "C-"
	} else if staff.RoleID == 2 {
		prefix = "S-"
	} else if staff.RoleID == 3 {
		prefix = "D-"
	} else {
		return "", errors.New("unexpected roleId")
	}

	staff.Code = prefix + util.GetRandomString(8)
	return staff.Code, m.db.CreateAccount(staff)
}

func (m *StaffManager) RemoveAccount(staffid int) error {
	return m.db.SetAccounQuit(staffid)	
}
