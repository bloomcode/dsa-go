// Package check is a tiny assertion helper shared by the tests.
package check

import (
	"reflect"
	"testing"
)

// Eq fails the test when got and want differ (deep comparison).
func Eq(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

// True fails the test when cond is false.
func True(t *testing.T, cond bool, msg string) {
	t.Helper()
	if !cond {
		t.Errorf("expected true: %s", msg)
	}
}
