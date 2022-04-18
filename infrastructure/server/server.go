package server

import "github.com/gin-gonic/gin"

var (
	engineGin = gin.New
	run       = startServer
)

func NewServer(server *gin.engine) {
	//definir si es dev || prod || QA
	server.Run(":3000")
}

func startServer() *gin.Engine {
	//Create server
	serverHTTP := engineGin()
	serverHTTP.Use(gin.Recovery())

	//set routes
	run(serverHTTP)
	return serverHTTP
}

setRoutes(router *gin.Engine){
	router.GET("/health",fmt.Println("health") )
	
}

