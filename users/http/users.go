package users

import (
	"fmt"
	"net/http"
	"sabasy/internal/dto"
	domain "sabasy/users/domain"
	users "sabasy/users/ports"

	"github.com/gin-gonic/gin"
)

type userRepositories struct {
	userRepos users.UserRepositories
}

/*users functions*/
func UserCreate(c *gin.Context) {

	var requestBody domain.User
	var response dto.Request

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
	}

	/*c.JSON(http.StatusOK, gin.H{
		"hello": c.Query("id"),
		"test":  c.PostForm("test"),
	})*/

	response.Data = requestBody
	response.Status = "200"

	c.JSON(http.StatusOK, response) //requestBody)
}

/*func UserUpdate()
func UserDelete()

func UserGetByID()
func UserGetAll()
func UserGetByFilterData()*/
