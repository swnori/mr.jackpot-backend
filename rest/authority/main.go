package authority



type AuthService interface {
	SigningService
	AuthMiddlewareService
}

type AuthHandler struct {
	SigningHandler
	AuthMiddlewareHandler
}