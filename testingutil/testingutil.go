package testingutil

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

const crasher = "CRASHER"
const crasherPlusOne = crasher + "=1"

// EvaluatePanic - executes the function and returns false if the function does not panic.
// Must be called from a top level Test method to ensure the test name is correct
func EvaluatePanic(t *testing.T, f func()) bool {
	t.Helper()

	if os.Getenv(crasher) == "1" {
		// function should panic
		f()
		return false
	}

	pcbuf := make([]uintptr, 1)
	n := runtime.Callers(2, pcbuf)
	callerName := ""
	if n != 0 {
		caller := runtime.FuncForPC(pcbuf[0] - 1)
		if caller != nil {
			tokens := strings.Split(caller.Name(), ".")
			callerName = tokens[len(tokens)-1]
		}
	}

	cmd := exec.Command(os.Args[0], "-test.run="+callerName)
	cmd.Env = append(os.Environ(), crasherPlusOne)
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)
	if ok && !e.Success() {
		return true
	}

	return false
}

// AssertCorrectString - tests got:want and throws t.Error with message in it if not ==
func AssertCorrectString(message string, t *testing.T, got, want string) {
	t.Helper()

	pcbuf := make([]uintptr, 1)
	n := runtime.Callers(2, pcbuf)
	callerName := ""
	if n != 0 {
		caller := runtime.FuncForPC(pcbuf[0] - 1)
		if caller != nil {
			tokens := strings.Split(caller.Name(), ".")
			callerName = tokens[len(tokens)-1]
		}
	}

	if got != want {
		if message == "" {
			t.Errorf("%v: got '%s' want '%s'", message, got, want)
			return
		}
		t.Errorf("%v/%v: got '%s' want '%s'", callerName, message, got, want)
	}
}

// AssertBoolean - tests ok and throws t.Error with message in it if not ==
func AssertBoolean(message string, t *testing.T, got, want bool) {
	t.Helper()

	pcbuf := make([]uintptr, 1)
	n := runtime.Callers(2, pcbuf)
	callerName := ""
	if n != 0 {
		caller := runtime.FuncForPC(pcbuf[0] - 1)
		if caller != nil {
			tokens := strings.Split(caller.Name(), ".")
			callerName = tokens[len(tokens)-1]
		}
	}

	if message == "" {
		t.Errorf("%v: got '%v' want '%v'", message, got, want)
		return
	}
	t.Errorf("%v/%v: got '%v' want '%v'", callerName, message, got, want)
}
