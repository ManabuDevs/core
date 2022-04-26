package main

import (
	dbconfig "sabasy/infrastructure/database"

	"fmt"
	"sabasy/infrastructure/server"
)

var start = server.StartServer

func main() {
	dbconfig.CreateDatabase()
	a := dbconfig.InstanceDB().GetConnect()
	fmt.Println(a)

	start().Run()
}
