package db

import (
	"context"

	"mr.jackpot-backend/model"
)




type StaffLayer interface {
	GetStaffRole(staffid int) (string, error)
	GetStaffID(staffID string) (int, error)
	CreateAccount(staff model.StaffInfo) error
	SetAccounQuit(staffid int) error

	GetStaffInfo(staffid int) (model.StaffInfo, error)
	GetAllStaffInfo() ([]model.StaffInfo, error)
}

type StaffDB struct {
	DBAccessor
}

func NewStaffDB() *StaffDB {
	db := &StaffDB{}
	db.q = NewAccessor()

	return db
}

func (db *StaffDB) GetStaffRole(staffid int) (string, error) {
	ctx := context.Background()
	return db.q.GetStaffRole(ctx, int64(staffid))
}