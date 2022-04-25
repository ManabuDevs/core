package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*users functions*/
func UserCreate(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

/*func UserUpdate()
func UserDelete()

func UserGetByID()
func UserGetAll()
func UserGetByFilterData()*/
