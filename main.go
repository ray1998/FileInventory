package main

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ray1998/workspaces/fileinventory/mssql"
)

var (
	// Trace provides trace logging
	Trace *log.Logger
	// Info provides info logging
	Info *log.Logger
	// Warning provides warning logging
	Warning *log.Logger
	// Error provides error logging
	Error *log.Logger
)

func main() {
	fmt.Println("File Inventory")

	connectionString := mssql.GetTrustedConnectionString("marble", "fileinventory", 0)
	db, err := mssql.OpenSQLTrustedConnection(connectionString)
	if err != nil {
		panic(err)
	}
	Trace.Println(db)

}
