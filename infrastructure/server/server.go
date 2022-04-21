package server

import (
	users "sabasy/users/routes"

	"github.com/gin-gonic/gin"
)

var (
	engineGin = gin.New
	run       = NewServer
)

func NewServer(server *gin.Engine) {
	//definir si es dev || prod || QA
	server.Run(":3000")
}

func StartServer() *gin.Engine {
	//Create server
	serverHTTP := engineGin()
	serverHTTP.Use(gin.Recovery())

	//set routes

	users.Routes(serverHTTP)
	run(serverHTTP)
	return serverHTTP
}
