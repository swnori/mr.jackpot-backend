package manager

import (
	"errors"
	"mr.jackpot-backend/model"
)

type CustomerAuthService interface {
	AuthService

	CheckAuthority(user model.User) (int, error)
	CreateAccound(user model.CustomerRegister) error
	RemoveAccount(userid int) error
}



func (m *CustomerManager) ComparePassword(s, t string) error {
	if s != t {
		return errors.New("Pass Not Match")
	}
	return nil
}

func (m *CustomerManager) CheckAuthority(user model.User) (int, error) {
	password, userid, err := m.db.GetCustomerPassword(user.UserID)
	if err != nil {
		return 0, err
	}
	if userid == 0 {
		return 0, errors.New("NO ID")
	}

	return userid, m.ComparePassword(password, user.Password);
}

func (m *CustomerManager) CreateAccound(user model.CustomerRegister) error {
	return m.db.CreateAccount(user);
}

func (m *CustomerManager) RemoveAccount(userid int) error {
	return m.db.SetAccounQuit(userid)
}
