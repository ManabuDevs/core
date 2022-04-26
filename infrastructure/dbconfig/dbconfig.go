package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	once     sync.Once
	instance *ConnectionDB
)

type ConnectionDB struct {
	connect *sql.DB
}

func ConnectDB(driveName string, urlConnect string) (*sql.DB, error) {
	return sql.Open(driveName, urlConnect)
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}

func (r ConnectionDB) GetConnect() *sql.DB {
	return r.connect
}

func InstanceDB() *ConnectionDB {
	once.Do(func() {
		connect, err := ConnectDB("sqlite3", "./test.db")
		if err != nil {
			panic(fmt.Sprintf("Don´t connection with DB: [%v]", err))
		}

		connect.SetMaxOpenConns(10)
		connect.SetMaxIdleConns(10)

		instance = &ConnectionDB{
			connect: connect,
		}
	})
	return instance
}

func CreateDatabase() {
	if _, err := os.Stat("./test.db"); os.IsNotExist(err) {
		file, err := os.Create("test.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
	}
}

/*
file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")

		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
		defer sqliteDatabase.Close()
*/
