package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/test", FindAll)
	}
}

func FindAll(c *gin.Context) {

	c.JSON(http.StatusOK, "{}")

}
