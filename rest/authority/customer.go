package authority

import "mr.jackpot-backend/service/manager"




type CustomerAuthService interface {
	RegisterService
	SigningService
}

type CustomerAuthHandler struct {
	m manager.CustomerService
}

func NewCustomerAuthHandler() *CustomerAuthHandler {
	return &CustomerAuthHandler{
		m: manager.NewCustomerManager(),
	}
}