package authority

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)



type CheckAuthMiddlewareService interface {
	CheckCEO(c *gin.Context)
	CheckCustomer(c *gin.Context)
	CheckCustomerOnly(c *gin.Context)
	CheckStaff(c *gin.Context)
}



func (h *AuthMiddlewareHandler) CheckCEO(c *gin.Context) {
	status := c.Keys["status"].(string)

	if status == model.UserStatusCEO {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, "Not a CEO")
	c.Abort()
}

func (h *AuthMiddlewareHandler) CheckCustomer(c *gin.Context) {
	status := c.Keys["status"].(string)

	if status == model.UserStatusCustomer || status == model.UserStatusVisitor {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, "Not a Customer")
	c.Abort()
}

func (h *AuthMiddlewareHandler) CheckCustomerOnly(c *gin.Context) {
	status := c.Keys["status"].(string)

	if status == model.UserStatusCustomer {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, "Not a Customer")
	c.Abort()
}
func (h *AuthMiddlewareHandler) CheckStaff(c *gin.Context) {
	status := c.Keys["status"].(string)

	if status == model.UserStatusStaff || status == model.UserStatusCEO {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, "Not a Staff")
	c.Abort()
}