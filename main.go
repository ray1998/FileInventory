package main

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ray1998/workspaces/fileinventory/mssql"
)

func main() {
	fmt.Println("Hello world")

	db, err := mssql.OpenSQLTrustedConnection("marble", "fileinventory")
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

}
