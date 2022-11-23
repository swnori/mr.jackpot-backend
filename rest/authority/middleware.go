package authority

import "mr.jackpot-backend/service/manager"



type AuthMiddlewareService interface {
	SetAuthMiddlewareService
	CheckAuthMiddlewareService
}

type AuthMiddlewareHandler struct {
	token manager.TokenService
}

func NewAuthMiddlewareHandler() *AuthMiddlewareHandler {
	return &AuthMiddlewareHandler{
		token: manager.DefaultToken,
	}
}