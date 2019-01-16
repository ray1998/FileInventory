package mssql

import (
	"os"
	"os/exec"
	"testing"
)

const Crasher = "CRASHER"

func TestGetTrustedConnectionString(t *testing.T) {

	t.Run("valid connection string", func(t *testing.T) {
		got := GetTrustedConnectionString("a", "b", 0)
		want := "server=a;database=b;Trusted_Connection=True;"
		assertCorrectMessage(t, got, want)
	})
}

func TestInvalidServer(t *testing.T) {
	if os.Getenv(Crasher) == "1" {
		// no server should panic
		GetTrustedConnectionString("", "b", 0)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidServer")
	cmd.Env = append(os.Environ(), Crasher+"=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process rn with err %v, want exit status 1", err)
}

func TestInvalidDatabase(t *testing.T) {
	if os.Getenv(Crasher) == "1" {
		// no database, should panic
		GetTrustedConnectionString("a", "", 0)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidDatabase")
	cmd.Env = append(os.Environ(), Crasher+"=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process rn with err %v, want exit status 1", err)
}
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("TestGetTrustedConnectionString: got '%s' want '%s'", got, want)
	}
}
