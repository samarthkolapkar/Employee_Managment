package connection

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func IntializeDB(conn string) {
	var err error
	db, err = sql.Open("postgres", conn)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error pinging the database: %v", err))
	}

	fmt.Println("Successfully connected to the database!")
}
func GetDB() *sql.DB {
	return db
}
