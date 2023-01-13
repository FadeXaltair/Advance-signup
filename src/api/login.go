package api

import (
	"login-crud/config"
	"login-crud/database"
	"login-crud/src/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login function is used to login the user
func Login(c *gin.Context) {
	var data database.Login
	err := c.Bind(&data)
	if err != nil {
		config.Error(err)
		return
	}

	var user database.Users
	res := config.DB.First(&user, "email= ?", data.Email)
	if res.Error != nil {
		config.Error(err)
		return
	}
	status := auth.CheckPasswordHash(data.Password, user.Password)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect password",
		})
		return
	}
	resp := database.Response{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  resp,
	})
}
