package main

import "testing"

func Test(t *testing.T) {
	if 1 == 0 {
		t.Fatal("why?")
	}
}
