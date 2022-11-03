package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	r := gin.Default()
	r.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})


	return r.Run(address)
}
