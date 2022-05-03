package misco_test

import (
	"testing"

	"github.io/jianzzha/hello/misco"
)

func TestMisco(t *testing.T) {
	if misco.BestMisco() != "cow" {
		t.Fatal("incorrect misco")
	}
}
