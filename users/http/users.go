package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*users functions*/
func UserCreate(c *gin.Context) {

	/*c.JSON(http.StatusOK, gin.H{
		"hello": c.Query("id"),
		"test":  c.PostForm("test"),
	})*/
	c.JSON(http.StatusOK, c.Request.Body)
}

/*func UserUpdate()
func UserDelete()

func UserGetByID()
func UserGetAll()
func UserGetByFilterData()*/
