package db

import (
	"context"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)




type StaffLayer interface {
	GetStaffID(code string) (int, error)
	CreateAccount(staff model.StaffRegister) error
	SetAccounQuit(staffid int) error

	GetStaffRole(staffid int) (string, error)
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

func (db *StaffDB) GetStaffID(code string) (int, error) {
	ctx := context.Background()

	staffid, err := db.q.GetStaffID(ctx, code)
	return int(staffid), err
}

func (db *StaffDB) CreateAccount(staff model.StaffRegister) error {
	ctx := context.Background()

	result, err :=db.q.CreateStaffAccount(ctx, orm.CreateStaffAccountParams{
		Name: staff.Name,
		RoleID: int32(staff.RoleID),
	})
	if err != nil {
		return err
	}

	staffid, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return db.q.CreateStaffAuth(ctx, orm.CreateStaffAuthParams{
		StaffID: staffid,
		Code: staff.Code,
	})
}

func (db *StaffDB) SetAccounQuit(staffid int) error {
	ctx := context.Background()
	return db.q.SetStaffQuit(ctx, int64(staffid))
}

func (db *StaffDB) GetStaffInfo(staffid int) (model.StaffInfo, error) {
	ctx := context.Background()

	staff, err := db.q.GetStaffInfo(ctx, int64(staffid))
	return model.StaffInfo{
		Name: staff.Name,
		Role: staff.Tag,
		Score: int(staff.Score),
	}, err
}

func (db *StaffDB) GetAllStaffInfo() ([]model.StaffInfo, error) {
	ctx := context.Background()

	staffList, err := db.q.GetAllStaffInfo(ctx)
	StaffList := make([]model.StaffInfo, 0)

	for _, staff := range staffList {
		StaffList = append(StaffList, model.StaffInfo{
			ID: int(staff.StaffID),
			Name: staff.Name,
			Role: staff.Tag,
			Score: int(staff.Score),
		})
	}

	return StaffList, err
}



/*
GetStaffID(staffID string) (int, error)
CreateAccount(staff model.StaffInfo) error
SetAccounQuit(staffid int) error

GetStaffInfo(staffid int) (model.StaffInfo, error)
GetAllStaffInfo() ([]model.StaffInfo, error)
*/