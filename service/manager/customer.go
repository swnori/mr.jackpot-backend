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

var Customer = &CustomerManager{}
func NewCustomerManager() *CustomerManager {
	Customer.Token = DefaultToken
	Customer.db = db.NewCustomerDB()
	return Customer
}