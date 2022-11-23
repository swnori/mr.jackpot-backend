package authority

import "mr.jackpot-backend/service/manager"




type VisitorAuthService interface {
	SigningService
}

type VisitorAuthHandler struct {
	m manager.VisitorAuthService
}

func NewVisitorAuthHandler() *VisitorAuthHandler {
	return &VisitorAuthHandler{
		m: manager.NewVisitorManager(),
	}
}