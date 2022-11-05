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
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.POST("/helloworld", func(c *gin.Context) {
		c.SetCookie("message",  "Hello World", 0,  "/", "https://mr-jackpot-backend.run.goorm.io", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Hello World Cookie Added",
		})
	})

	return r.Run(address)
}
