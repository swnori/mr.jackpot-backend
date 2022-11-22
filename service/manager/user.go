package manager

import "errors"


type Manager struct {}

var errMessage string = "Implement Inherited Class"

func (m *Manager) CheckAuthority() error {
	return errors.New(errMessage)
}
func (m *Manager) CreateAccound() error {
	return errors.New(errMessage)
}
func (m *Manager) RemoveAccount() error {
	return errors.New(errMessage)
}

func (m *Manager) GetUserInfo() error {
	return errors.New(errMessage)
}
func (m *Manager) GetAllUserInfo() error {
	return errors.New(errMessage)
}
func (m *Manager) UpdateUserInfo() error {
	return errors.New(errMessage)
}
