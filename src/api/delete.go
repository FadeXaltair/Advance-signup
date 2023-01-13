package api

import (
	"log"
	"login-crud/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete function to delete the data from the database
func Delete(c *gin.Context) {
	userid, _ := c.GetQuery("user-id")
	id, _ := strconv.Atoi(userid)
	err := DeleteData(id)
	if err != nil {
		log.Println(err)
		config.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Account deleted successfully",
	})
}

func DeleteData(id int) error {
	err := config.DB.Exec(`delete from public.users  where id =` + strconv.Itoa(id) + ``).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
