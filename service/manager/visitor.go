package manager

import (

	"mr.jackpot-backend/database/db"
)



type VisitorService interface {
	AuthService
}

type VisitorManager struct {
	db db.VisitorLayer
	Client
}

var Visitor = &VisitorManager{}

func NewVisitorManager() *VisitorManager {
	Visitor.Token = DefaultToken
	Visitor.db = db.NewVisitorDB()
	return Visitor
}