package manager

import "mr.jackpot-backend/database/db"



type CustomerService interface {
	CustomerAuthService
	CustomerInfoService
}

type CustomerManager struct {
	db db.CustomerLayer

	Client
	Member
}

var Customer = &CustomerManager{
	db: db.NewCustomerDB(),
}

func NewCustomerManager() *CustomerManager {
	Customer.Token = DefaultToken
	return Customer
}