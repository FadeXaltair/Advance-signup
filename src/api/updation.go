package api

import (
	"log"
	"login-crud/config"
	"login-crud/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Update function is used to update the data of user
func Update(c *gin.Context) {
	userid, _ := c.GetQuery("user-id")
	id, _ := strconv.Atoi(userid)
	var data database.Updations
	err := c.Bind(&data)
	if err != nil {
		log.Println(err)
		config.Error(err)
	}
	err = UpdateData(id, data.Name, data.Email)
	if err != nil {
		log.Println(err)
		config.Error(err)
	}
	resp := database.Response{
		Id:    id,
		Name:  data.Name,
		Email: data.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  resp,
	})
}

func UpdateData(id int, name, email string) error {
	err := config.DB.Exec(`update public.users 
	set name= '` + name + `' , email ='` + email + `'
	where id =` + strconv.Itoa(id) + ``).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
