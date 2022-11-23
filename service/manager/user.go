package manager

import "errors"



type Member struct {}

func NewMember() *Member {
	return &Member{}
}



var errMessage string = "Implementing Inherited Class Required"

func (m *Member) CheckAuthority() error {
	return errors.New(errMessage)
}
func (m *Member) CreateAccound() error {
	return errors.New(errMessage)
}
func (m *Member) RemoveAccount() error {
	return errors.New(errMessage)
}

func (m *Member) GetUserInfo() error {
	return errors.New(errMessage)
}
func (m *Member) GetAllUserInfo() error {
	return errors.New(errMessage)
}
func (m *Member) UpdateUserInfo() error {
	return errors.New(errMessage)
}

