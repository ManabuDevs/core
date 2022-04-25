package main

import (

	//dbconfig "sabasy/infrastructure/database"
	"database/sql"
	"log"
	"os"
	"sabasy/infrastructure/server"

	_ "github.com/mattn/go-sqlite3"
)

var start = server.StartServer

func main() {
	//a := dbconfig.InstanceDB().GetConnect()
	//fmt.Println(a)
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()

	start().Run()
}
