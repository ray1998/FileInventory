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

// ConnectionString defines a sql connection string
type ConnectionString struct {
	Server            string
	Database          string
	TrustedConnection bool
	Port              int
	User              string
	Password          string
}

// OpenSQLTrustedConnection opens a sql database
func OpenSQLTrustedConnection(server, database string) (*sql.DB, error) {
	// Build connection string
	connString := fmt.Sprintf("server=%s;database=%s;Trusted_Connection=True;", server, database)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return nil, err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	fmt.Printf("Connected to server: %s database: %s\n", server, database)

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
