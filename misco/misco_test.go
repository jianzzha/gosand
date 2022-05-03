package misco_test

import (
	"testing"

	"github.com/jianzzha/gosand/misco"
)

func TestMisco(t *testing.T) {
	if misco.BestMisco() != "cow" {
		t.Fatal("incorrect misco")
	}
}
