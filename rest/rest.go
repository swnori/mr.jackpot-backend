package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func RunAPI(address string) error {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://mr-jackpot.run.goorm.io/"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.POST("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	return r.Run(address)
}
