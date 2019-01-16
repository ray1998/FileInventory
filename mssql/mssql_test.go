package mssql_test

import (
	"testing"

	mssql "github.com/ray1998/workspaces/FileInventory/mssql"
	testingutil "github.com/ray1998/workspaces/FileInventory/testingutil"
)

func TestGetTrustedConnectionString(t *testing.T) {
	got := mssql.GetTrustedConnectionString("a", "b", 0)
	want := "server=a;database=b;Trusted_Connection=True;"
	testingutil.AssertCorrectString("Invalid connection string", t, got, want)
}

func TestInvalidServer(t *testing.T) {
	ok := testingutil.EvaluatePanic(t, func() {
		mssql.GetTrustedConnectionString("", "b", 0)
	})

	testingutil.AssertBoolean("Did not panic", t, ok, true)
}

func TestInvalidDatabase(t *testing.T) {
	ok := testingutil.EvaluatePanic(t, func() {
		mssql.GetTrustedConnectionString("a", "", 0)
	})

	testingutil.AssertBoolean("Did not panic", t, ok, true)
}
