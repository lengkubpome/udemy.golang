package main

import "testing"

func TestReadLines(t *testing.T) {
	want := []string{"a", "b", "x"}
	got := readLines("./test/todos.txt")
	if !equalSlice(want, got) {
		t.Errorf("readLines(\"./test/todos.txt\") =%q, want %q", got, want)
	}

}

func equalSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}
