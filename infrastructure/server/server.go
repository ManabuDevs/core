package server

import "github.com/gin-gonic/gin"

var (
	engineGin = gin.New
	run       = NewServer
)

func NewServer(server *gin.Engine) {
	//definir si es dev || prod || QA
	server.Run(":3000")
}

func StartServer(routes []string) *gin.Engine {
	//Create server
	serverHTTP := engineGin()
	serverHTTP.Use(gin.Recovery())

	//set routes
	setRoutes(serverHTTP)
	run(serverHTTP)
	return serverHTTP
}

setRoutes(router *gin.Engine){
	router.GET("/health",fmt.Println("health") )
}
