package routes

import (
	"login-crud/src/api"

	"github.com/gin-gonic/gin"
)

// Routes function is used for the routing for our apis
func Routes(router *gin.Engine) {
	router.POST("/signup", api.SignUp)
	router.POST("/login", api.Login)
	router.PUT("/update", api.Update)
	router.GET("/users", api.FilterUser)
	router.GET("/details", api.Details)
	router.DELETE("/users", api.Delete)
	router.Run()
}
