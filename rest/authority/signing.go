package authority

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)



type SigningService interface {
	Signin(c *gin.Context)
	Signout(c *gin.Context)
}

func (h *CustomerAuthHandler) Signin(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	userid, err := h.m.CheckAuthority(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.m.CreateToken(userid, model.UserStatusCustomer)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.SetCookie("access-token", token, 0, "/", "http://127.0.0.1:3000", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access-token": token,
	})
}

func (h *CustomerAuthHandler) Signout(c *gin.Context) {
	c.SetCookie("access-token",  "", 0, "/", "http://127.0.0.1:3000", false, true)
	c.JSON(http.StatusOK, "")
}



func (h *StaffAuthHandler) Signin(c *gin.Context) {

}

func (h *StaffAuthHandler) Signout(c *gin.Context) {
	c.SetCookie("access-token",  "", 0, "/", "http://127.0.0.1:3000", false, true)
	c.JSON(http.StatusOK, "")
}



func (h *VisitorAuthHandler) Signin(c *gin.Context) {
	var (
		err error
		identifier string
		userid int
		tokenString string
	)

	identifier, err = c.Cookie("identifier")
	if err != nil {
		identifier, err = h.Register()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	userid, err = h.m.CheckAuthority(identifier)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	tokenString, err = h.m.CreateToken(userid, model.UserStatusVisitor)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("access-token", tokenString, 0, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("identifier",   identifier,  0, "/", "http://127.0.0.1:3000", false, true)
	c.JSON(http.StatusOK, "")
}



func (h *VisitorAuthHandler) Signout(c *gin.Context) {
	c.SetCookie("access-token",  "", 0, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("identifier",  "", 0, "/", "http://127.0.0.1:3000", false, true)
	c.JSON(http.StatusOK, "")
}
