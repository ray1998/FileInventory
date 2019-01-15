package mssql

import (
	"context"
	"database/sql"
	"log"

	// "context"
	// "log"
	"fmt"
	// "errors"
)

// // ConnectionString defines a sql connection string
// type ConnectionString struct {
// 	Server            string
// 	Database          string
// 	TrustedConnection bool
// 	Port              int
// 	User              string
// 	Password          string
// }

// GetTrustedConnectionString - creates a trusted sql connection string
func GetTrustedConnectionString(server, database string, port int) string {
	if server == "" || database == "" {
		msg := fmt.Sprintf("[server=%s;database=%s;Trusted_Connection=True;] is invalid", server, database)
		log.Fatal(msg)
	}

	if port > 0 {
		return fmt.Sprintf("server=%s:%d;database=%s;Trusted_Connection=True;", server, port, database)

	}

	return fmt.Sprintf("server=%s;database=%s;Trusted_Connection=True;", server, database)
}

// OpenSQLTrustedConnection opens a sql database
func OpenSQLTrustedConnection(connectionString string) (*sql.DB, error) {

	// Create connection pool
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db, nil
}

// Reverse will reverse a string
// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}

// 	val, err := mssqldb.Float64ToDecimal(3.1)
// 	fmt.Println(val, err)
// 	return string(runes)
// }
