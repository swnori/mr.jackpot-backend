package manager

import (
	"errors"
	"mr.jackpot-backend/model"
)

type CustomerAuthService interface {
	CheckAuthority(user model.User) error
	CreateAccound(user model.RegisterRequest) error
	RemoveAccount(userid int) error
}



func (m *CustomerManager) ComparePassword(s, t string) error {
	if s != t {
		return errors.New("Pass Not Match")
	}
	return nil
}

func (m *CustomerManager) CheckAuthority(user model.User) error {
	if err := m.db.CheckUserExist(user.UserID); err != nil {
		return err
	}
	
	password, err := m.db.GetUserPassword(user.UserID)
	if err != nil {
		return err
	}

	return m.ComparePassword(password, user.Password);
}

func (m *CustomerManager) CreateAccound(user model.RegisterRequest) error {
	return m.db.CreateAccount(user);
}

func (m *CustomerManager) RemoveAccount(userid int) error {
	return m.db.SetAccounQuit(userid)
}
