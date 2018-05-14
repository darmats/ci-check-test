package check

import (
	"testing"
)

func TestVet(t *testing.T) {
	if !Vet() {
		t.Error("want true, got false")
	}
}
