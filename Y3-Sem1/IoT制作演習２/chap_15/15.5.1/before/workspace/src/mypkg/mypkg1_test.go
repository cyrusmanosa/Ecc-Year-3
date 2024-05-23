package mypkg_test

import (
	"mypkg"
	"testing"
)

func TestAdd(t *testing.T) {
	if mypkg.Add(1, 2) != 3 {
		t.Fail()
	}
}
