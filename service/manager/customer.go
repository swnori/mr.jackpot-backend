package manager

import "mr.jackpot-backend/database/db"



type CustomerService interface {
	CustomerAuthService
	CustomerInfoService
}

type CustomerManager struct {
	db db.CustomerLayer
	Manager
}

var Customer = &CustomerManager{
	db: db.NewCustomerDB(),
}

func NewCustomerManager() *CustomerManager {
	return Customer
}