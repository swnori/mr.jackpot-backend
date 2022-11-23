package authority

import (
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)


type SetAuthMiddlewareService interface {
	SetAuthority(c *gin.Context)
}



func (h *AuthMiddlewareHandler) SetAuthority(c *gin.Context) {
	tokenString, err := c.Cookie("access-token")
	if err != nil {
		c.Set("status", model.UserStatusUnauthorized)
		c.Set("userid", 0)
		c.Next()
		return
	}

	userid, status, err := h.token.ParseToken(tokenString)
	if err != nil {
		c.Set("status", model.UserStatusUnauthorized)
		c.Set("userid", 0)
		c.Next()
		return
	}

	c.Set("status", status)
	c.Set("userid", userid)
	c.Next()
}

