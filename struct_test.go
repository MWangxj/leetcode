package main

import (
	"math/rand"
	"testing"
)

type st struct {
}

func TestStruct(t *testing.T) {
	s := &st{}
	if s == nil {
		t.Logf("fail")
	}
	t.Logf("pass")
}

func TestStruct2(t *testing.T) {
	a := rand.Int63n(10)
	if a > 5 {
		t.Fatalf("fail")
	}
	t.Logf("success\n")
}
