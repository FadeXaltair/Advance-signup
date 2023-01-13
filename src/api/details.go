package api

import (
	"log"
	"login-crud/config"
	"login-crud/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Details function will give the details of the user
func Details(c *gin.Context) {
	userid, _ := c.GetQuery("user-id")
	id, _ := strconv.Atoi(userid)
	data, err := Data(id)
	if err != nil {
		log.Println(err)
		config.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}

func Data(id int) (database.Details, error) {
	var data database.Details
	err := config.DB.Raw("select * from public.users u where u.id =" + strconv.Itoa(id) + "").Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}
