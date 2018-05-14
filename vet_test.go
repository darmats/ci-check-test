package check_test

import (
	"testing"
	"github.com/darmats/circleci-check-test"
)

func TestVet(t *testing.T) {
	if !check.Vet() {
		t.Error("want true, got false")
	}
}
