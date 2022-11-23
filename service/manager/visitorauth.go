package manager


type VisitorAuthService interface {
	AuthService

	CheckAuthority(identifier string) (int, error)
	CreateAccount(identifier string) (int, error)
}


func (m *VisitorManager) CheckAuthority(identifier string) (int, error) {
	return m.db.GetVisitorByIdentifier(identifier)
}

func (m *VisitorManager) CreateAccount(identifier string) (int, error) {
	return m.db.CreateVisitor(identifier)
}