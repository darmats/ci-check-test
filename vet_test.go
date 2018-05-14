package check_test

import (
	"testing"
	"github.com/darmats/ci-check-test"
)

func TestVet(t *testing.T) {
	if !check.Vet() {
		t.Error("want true, got false")
	}
}
