package main

import (
	"sabasy/infrastructure/server"
)

var start = server.StartServer

func main() {
	start([]string{"Crimson", "Red", "Ruby", "Maroon"}).Run()
}
