package users

import (
	users "sabasy/users/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.POST("/create", users.UserCreate)
		user.GET("/all", users.UserGetAll)
		user.GET("/id/:user_id", users.UserGetByID)
		user.DELETE("/id/:user_id", users.UserDeleteByID)
	}
}
