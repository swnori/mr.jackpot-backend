package db

import "mr.jackpot-backend/model"




type CustomerLayer interface {
	CheckUserExist(userid string) error
	GetUserPassword(userid string) (string, error)
	GetUserAuth() (model.User, error)

	CreateAccount(customer model.RegisterRequest) error
	SetAccounQuit(userid int) error

	GetCustomerInfo(userid int) (model.CustomerInfo, error)
	GetAllCustomerInfo() ([]model.CustomerInfo, error)

	GetPersonalnfo(userid int) (model.PersonalInfo, error)
	UpdateUserInfo(userid int, personal model.PersonalInfo) error
}


type CustomerDB struct {
	DBAccessor
}

func NewCustomerDB() *CustomerDB {
	db := &CustomerDB{}
	db.q = NewAccessor()

	return db
}