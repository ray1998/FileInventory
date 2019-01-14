package main

import (
	"fmt"
	// the sql functions
	"github.com/ray1998/workspaces/FileInventory/mssql"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(mssql.Reverse("hello"))
}
