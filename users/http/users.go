package users

import (
	"fmt"
	"net/http"
	"sabasy/infrastructure/dbconfig"
	"sabasy/internal/dto"
	domain "sabasy/users/domain"
	usersRepo "sabasy/users/repo"
	usersService "sabasy/users/services"

	"github.com/gin-gonic/gin"
)

/*users functions*/
func UserGetAll(c *gin.Context) {
	var response dto.Request

	repo := usersRepo.NewUserRepository(dbconfig.InstanceDB().GetConnect())
	service := usersService.NewUserService(repo)

	t, _ := service.GetUsers()

	response.Data = t
	response.Status = "200"

	c.JSON(http.StatusOK, response)
}

func UserGetByID(c *gin.Context) {
	var response dto.Request

	repo := usersRepo.NewUserRepository(dbconfig.InstanceDB().GetConnect())
	service := usersService.NewUserService(repo)

	t, _ := service.GetUserByID(c.Param("user_id"))

	//pasar a utils esta funcion
	if len(t) > 0 {
		response.Data = t
		c.JSON(http.StatusOK, response)
	} else {
		response.Data = "no existe el id a buscar"
	}

	response.Status = "204" //esto tambien a utils con el de arriba
	c.JSON(http.StatusOK, response)
}

func UserCreate(c *gin.Context) {
	var requestBody domain.User
	var response dto.Request

	repo := usersRepo.NewUserRepository(dbconfig.InstanceDB().GetConnect())
	service := usersService.NewUserService(repo)

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
	}

	t, _ := service.CreateUser(&requestBody)
	//usersPorts.UserRepositories.CreateUser(requestBody)
	/*c.JSON(http.StatusOK, gin.H{
		"hello": c.Query("id"),
		"test":  c.PostForm("test"),
	})*/

	response.Data = t
	response.Status = "200"

	c.JSON(http.StatusOK, response) //requestBody)
}

func UserDeleteByID(c *gin.Context) {
	var response dto.Request

	repo := usersRepo.NewUserRepository(dbconfig.InstanceDB().GetConnect())
	service := usersService.NewUserService(repo)

	t, _ := service.DeleteUserByID(c.Param("user_id"))

	response.Data = t
	response.Status = "200" //esto tambien a utils con el de arriba
	c.JSON(http.StatusOK, response)
}

/*func UserUpdate()
func UserGetByFilterData()*/
