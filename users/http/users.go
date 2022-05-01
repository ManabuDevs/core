package users

import (
	"fmt"
	"net/http"
	domain "sabasy/users/domain"

	"github.com/gin-gonic/gin"
)

/*users functions*/
func UserCreate(c *gin.Context) {

	var requestBody domain.User

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
	}

	/*c.JSON(http.StatusOK, gin.H{
		"hello": c.Query("id"),
		"test":  c.PostForm("test"),
	})*/
	c.JSON(http.StatusOK, requestBody)
}

/*func UserUpdate()
func UserDelete()

func UserGetByID()
func UserGetAll()
func UserGetByFilterData()*/
