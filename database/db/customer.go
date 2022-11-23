package db

import (
	"context"
	"database/sql"
	"errors"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)




type CustomerLayer interface {
	GetCustomerPassword(id string) (string, int, error)
	CreateAccount(customer model.CustomerRegister) error
	SetAccounQuit(userid int) error

	GetPersonalnfo(userid int) (model.PersonalInfo, error)
	UpdateUserInfo(userid int, personal model.PersonalInfo) error
	GetAllCustomerInfo() ([]model.CustomerInfo, error)
}


type CustomerDB struct {
	DBAccessor
}

func NewCustomerDB() *CustomerDB {
	db := &CustomerDB{}
	db.q = NewAccessor()

	return db
}


func (db *CustomerDB) GetCustomerPassword(id string) (string, int, error) {
	ctx := context.Background()
	result, err := db.q.GetCustomerPassword(ctx, id)
	return result.Password, int(result.CustomerID), err
}

func (db *CustomerDB) CreateAccount(customer model.CustomerRegister) error {
	ctx := context.Background()

	flag, err := db.q.CheckCustomerID(ctx, customer.UserID)
	if flag != 0 {
		return errors.New("ID already exist")
	}

	result, err := db.q.CreateUser(ctx)
	if err != nil {
		return err
	}
	userid, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if err := db.q.CreateCustomer(ctx, orm.CreateCustomerParams{
		CustomerID: userid,
		Name: customer.Name,
		Address: sql.NullString{
			String: customer.Address,
			Valid: true,
		},
		Phone: sql.NullString{
			String: customer.Phone,
			Valid: true,
		},
	}); err != nil {
		return err
	}

	return db.q.CreateCustomerAuth(ctx, orm.CreateCustomerAuthParams{
		ID: customer.UserID,
		Password: customer.Password,
		CustomerID: userid,
	})
}

func (db *CustomerDB) SetAccounQuit(userid int) error {
	ctx := context.Background()
	return db.q.SetCustomerQuit(ctx, int64(userid))
}

func (db *CustomerDB) GetAllCustomerInfo() ([]model.CustomerInfo, error) {
	ctx := context.Background()

	customerList, err := db.q.GetAllCustomerInfo(ctx)
	
	CustomerList := make([]model.CustomerInfo, 0)
	for _, customer := range customerList {
		CustomerList = append(CustomerList, model.CustomerInfo{
			ID: int(customer.CustomerID),
		})
	}

	return CustomerList, err
}

func (db *CustomerDB) GetPersonalnfo(userid int) (model.PersonalInfo, error) {
	ctx := context.Background()

	personal, err := db.q.GetPersonalInfo(ctx, int64(userid))
	return model.PersonalInfo{
		ID: int(personal.CustomerID),
		Name: personal.Name,
		Address: personal.Address.String,
		Phone: personal.Phone.String,
	}, err
}

func (db *CustomerDB) UpdateUserInfo(userid int, personal model.PersonalInfo) error {
	ctx := context.Background()

	return db.q.UpdatePersonalInfo(ctx, orm.UpdatePersonalInfoParams{
		CustomerID: int64(personal.ID),
		Name: personal.Name,
		Phone: sql.NullString{
			String: personal.Phone,
			Valid: true,
		},
		Address: sql.NullString{
			String: personal.Address,
			Valid: true,
		},
	})
}