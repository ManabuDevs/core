package db

import (
	"database/sql"
	"fmt"
	"sync"
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
		connect, err := ConnectDB("sqlite", "./test")
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
