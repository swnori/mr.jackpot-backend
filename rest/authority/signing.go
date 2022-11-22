package authority

import "github.com/gin-gonic/gin"



type SigningService interface {
	Signin(c *gin.Context)
	Signout(c *gin.Context)
	LoginForCustomers(c *gin.Context)
	LogoutForCustomers(c *gin.Context)
	LoginForStaff(c *gin.Context)
	LogoutForStaff(c *gin.Context)	
}

type SigningHandler struct {

}


func (h *SigningHandler) Signin(c *gin.Context) {}
func (h *SigningHandler) Signout(c *gin.Context) {}

func (h *SigningHandler) LoginForCustomers(c *gin.Context) {}
func (h *SigningHandler) LogoutForCustomers(c *gin.Context) {}
func (h *SigningHandler) LoginForStaff(c *gin.Context) {}
func (h *SigningHandler) LogoutForStaff(c *gin.Context) {}




