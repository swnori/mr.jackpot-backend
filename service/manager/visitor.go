package manager

import "mr.jackpot-backend/database/db"



type VisitorService interface {
	AuthService
}

type VisitorManager struct {
	db db.VisitorLayer
	Client
}

var Visitor = &VisitorManager{
	db: db.NewVisitorDB(),
}

func NewVisitorManager() *VisitorManager {
	Visitor.Token = DefaultToken
	return Visitor
}