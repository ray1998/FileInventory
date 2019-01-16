package testingutil_test

import (
	"testing"

	testingutil "github.com/ray1998/workspaces/FileInventory/testingutil"
)

func TestPanicValid(t *testing.T) {
	ok := testingutil.EvaluatePanic(t, func() {
		panic("expected")
	})

	testingutil.AssertBoolean("did not panic", t, ok, true)
}

func TestPanicInvalid(t *testing.T) {
	ok := testingutil.EvaluatePanic(t, func() {
		// no panic
		return
	})

	testingutil.AssertBoolean("did not panic", t, ok, false)
}
