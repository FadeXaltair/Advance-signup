package api

import (
	"log"
	"login-crud/config"
	"login-crud/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Filteruser function is used to filter the data with the name of user and
// if name feild is empty, then it will show all users data
func FilterUser(c *gin.Context) {
	name, _ := c.GetQuery("name")
	if name != "" {
		data, err := AllData(name)
		if err != nil {
			log.Println(err)
			config.Error(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  data,
		})
		return
	}

	data, err := FullData()
	if err != nil {
		log.Println(err)
		config.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}

func AllData(name string) ([]database.Response, error) {
	var data []database.Response
	err := config.DB.Raw(`select u.id , u.name, u.email  from public.users u where name ilike '%` + name + `%'`).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func FullData() ([]database.Response, error) {
	var data []database.Response
	err := config.DB.Raw(`select u.id , u.name, u.email  from public.users u`).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}
