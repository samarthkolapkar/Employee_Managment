package connection

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func IntializeDB(conn string) {
	// Initialize db connection during package initialization
	// host := "localhost"
	// port := 5432
	// user := "postgres"
	// password := "Sbk@5885"
	// dbname := "Employee"
	// sslmode := "disable"

	// connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// host, port, user, password, dbname, sslmode)

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
