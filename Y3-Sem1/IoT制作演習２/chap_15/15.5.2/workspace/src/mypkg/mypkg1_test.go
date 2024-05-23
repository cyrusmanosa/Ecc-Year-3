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

// MapLoop1関数のベンチマーク
func BenchmarkMapLoop1(b *testing.B) {
	mypkg.MapLoop1()
}

// MapLoop2関数のベンチマーク
func BenchmarkMapLoop2(b *testing.B) {
	mypkg.MapLoop2()
}
