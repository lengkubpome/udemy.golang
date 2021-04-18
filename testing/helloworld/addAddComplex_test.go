package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := add(1, 2)
	if a != 3 {
		t.Error("a := add(1,2) : is not 3")
	}
}

func TestAdd2(t *testing.T) {
	a := add(2, 2)
	if a != 4 {
		t.Error("a := add(1,2) : is not 3")
	}
}
