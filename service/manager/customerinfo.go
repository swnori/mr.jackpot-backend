package manager

import "mr.jackpot-backend/model"

type CustomerInfoService interface {
	GetUserInfo(userid int) (model.PersonalInfo, error)
	GetAllUserInfo() ([]model.CustomerInfo, error)
	UpdateUserInfo(userid int, personal model.PersonalInfo) error
	GetPersonalnfo(userid int) (model.PersonalInfo, error)
}



func (m *CustomerManager) GetUserInfo(userid int) (model.PersonalInfo, error) {
	return m.db.GetPersonalnfo(userid)
}

func (m *CustomerManager) GetAllUserInfo() ([]model.CustomerInfo, error) {
	return m.db.GetAllCustomerInfo()
}

func (m *CustomerManager) UpdateUserInfo(userid int, personal model.PersonalInfo) error {
	return m.db.UpdateUserInfo(userid, personal)
}

func (m *CustomerManager) GetPersonalnfo(userid int) (model.PersonalInfo, error) {
	return m.db.GetPersonalnfo(userid)
}
