package main

import (
	"sabasy/infrastructure/dbconfig"
	"sabasy/infrastructure/server"
)

var (
	start    = server.StartServer
	createDB = dbconfig.CreateDatabase
)

func main() {

	createDB()
	start().Run()

}
