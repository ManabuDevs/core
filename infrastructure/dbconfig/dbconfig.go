package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
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

var valueStringNil string

func ParserQuery(rows *sql.Rows, dynamics map[string]string) (result []map[string]interface{}) {
	for rows.Next() {
		keys, _ := rows.Columns()
		cols := make([]interface{}, len(keys))
		for i := range keys {
			cols[i] = &cols[i]
		}
		rows.Scan(cols...)

		obj := make(map[string]interface{})
		for i, key := range keys {
			d := dynamics[key]
			if d != "" {
				obj[d] = ParseString(cols[i])
			}
		}
		result = append(result, obj)
	}
	return
}

func parseArrayUInt8ToString(data []uint8) string {
	if data != nil {
		var valueContent strings.Builder
		valueContent.Write(data)
		return valueContent.String()
	}
	return valueStringNil
}

func ParseString(data interface{}) string {
	if data != nil {
		switch data.(type) {
		case []uint8:
			return parseArrayUInt8ToString(data.([]uint8))
		default:
			return fmt.Sprintf("%v", data)
		}
	}
	return valueStringNil
}
