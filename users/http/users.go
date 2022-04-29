package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*users functions*/
func UserCreate(c *gin.Context) {
	var x []string
	c.JSON(http.StatusOK, x)
}

/*func UserUpdate()
func UserDelete()

func UserGetByID()
func UserGetAll()
func UserGetByFilterData()*/
