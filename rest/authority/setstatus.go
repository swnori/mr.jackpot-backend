package authority

import (
	"strings"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)


type SetAuthMiddlewareService interface {
	SetAuthority(c *gin.Context)
}



func (h *AuthMiddlewareHandler) SetAuthority(c *gin.Context) {
	tokenString, err := c.Cookie("access-token")
	if err != nil {
		bareToken := c.Request.Header["Authorization"]
		if len(bareToken) != 1{
			c.Set("status", model.UserStatusUnauthorized)
			c.Set("userid", 0)
			c.Next()
			return
		}
		tokenParsed := strings.Split(bareToken[0], " ")

		if len(tokenParsed) != 2  {
			c.Set("status", model.UserStatusUnauthorized)
			c.Set("userid", 0)
			c.Next()
			return
		}
		tokenString = tokenParsed[1]
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

