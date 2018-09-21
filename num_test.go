package main

import (
	`testing`
)

func TestCompare(t *testing.T) {
	a := 1<<20
	b :=1024 *1024

	if a == b {
		t.Log("a equal b")
		return
	}
	t.Fail()
}
