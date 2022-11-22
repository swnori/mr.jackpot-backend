package authority

import "github.com/gin-gonic/gin"


type AuthMiddlewareService interface {
	SetAuthority(c *gin.Context)
	SetAuthorityMiddleware() gin.HandlerFunc

	CheckCEO(c *gin.Context)
	CheckCEOMiddleware() gin.HandlerFunc
	CheckCustomer(c *gin.Context)
	CheckCustomerMiddleware() gin.HandlerFunc
	CheckStaff(c *gin.Context)
	CheckStaffMiddleware() gin.HandlerFunc
}

type AuthMiddlewareHandler struct {
	
}

func (h *AuthMiddlewareHandler) SetAuthority(c *gin.Context) {}
func (h *AuthMiddlewareHandler) SetAuthorityMiddleware() gin.HandlerFunc {
	return h.SetAuthority
}

func (h *AuthMiddlewareHandler) CheckCEO(c *gin.Context) {}
func (h *AuthMiddlewareHandler) CheckCEOMiddleware() gin.HandlerFunc {
	return h.CheckCEO
}

func (h *AuthMiddlewareHandler) CheckCustomer(c *gin.Context) {}
func (h *AuthMiddlewareHandler) CheckCustomerMiddleware() gin.HandlerFunc {
	return h.CheckCustomer
}

func (h *AuthMiddlewareHandler) CheckStaff(c *gin.Context) {}
func (h *AuthMiddlewareHandler) CheckStaffMiddleware() gin.HandlerFunc {
	return h.CheckStaff
}