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

func TestNoPanicShouldFail(t *testing.T) {
	ok := testingutil.EvaluatePanic(t, func() {
		// no panic
		return
	})

	testingutil.AssertBoolean("should fail", t, ok, true)
}

func TestAssertBooleanShouldFail(t *testing.T) {
	testingutil.AssertBoolean("should fail", t, true, false)
}

func TestAssertCorrectString(t *testing.T) {
	testingutil.AssertCorrectString("Should not fail", t, "a", "a")
}

func TestAssertCorrectStringShouldFail(t *testing.T) {
	testingutil.AssertCorrectString("Should fail", t, "a", "b")
}
