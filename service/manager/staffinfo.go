package manager

import (
	"errors"

	"mr.jackpot-backend/model"
)

type StaffInfoService interface {
	GetUserInfo(staffid int) (model.StaffInfo, error)
	GetAllUserInfo() ([]model.StaffInfo, error)
	UpdateUserInfo() error
}

func (m *StaffManager) GetUserInfo(staffid int) (model.StaffInfo, error) {
	return m.db.GetStaffInfo(staffid)
}

func (m *StaffManager) GetAllUserInfo() ([]model.StaffInfo, error) {
	return m.db.GetAllStaffInfo()
}

func (m *StaffManager) UpdateUserInfo() error {
	return errors.New("Service Not Provided")
}