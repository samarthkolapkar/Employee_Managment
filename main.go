package main

import (
	"employee/connection"
	"employee/handlers"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// http.HandleFunc("/addEmployee", handlers.HandleAddEmployee)
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "Employee"
	sslmode := "disable"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	connection.IntializeDB(dsn)
	http.HandleFunc("/employeeList", handlers.HandleEmployeeList)
	http.HandleFunc("/addEmployee", handlers.HandleAddEmployee)
	http.HandleFunc("/employee_get", handlers.HandleEmployeeForward)
	http.HandleFunc("/employee_approve", handlers.HandleEmployeeApprove)
	http.HandleFunc("/inactive_employee", handlers.HandleEmployeeInactive)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
