package manager

import "mr.jackpot-backend/service/manager"



type ManagerService interface {
	StaffManagerService
	CustomerManagerService
}

type ManagerHandler struct {
	sm manager.StaffService
	cm manager.CustomerService
}


func NewManagerHandler() *ManagerHandler {
	return &ManagerHandler{
		sm: manager.NewStaffManager(),
		cm: manager.NewCustomerManager(),
	}
}