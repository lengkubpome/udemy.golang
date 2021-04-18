package main

import (
	"testing"
)

func TestAddComplex(t *testing.T) {
	tests := []struct {
		lhs      int
		rhs      int
		expected int
	}{
		{lhs: 1, rhs: 2, expected: 4},
		{lhs: 2, rhs: 2, expected: 4},
		{lhs: 3, rhs: 2, expected: 5},
		{lhs: 10, rhs: 2, expected: 12},
		{lhs: 10, rhs: -5, expected: 15},
	}

	for _, test := range tests {
		ans := add(test.lhs, test.rhs)
		if ans != test.expected {
			t.Errorf("add(%d,%d) = %d . Wanted %d", test.lhs, test.rhs, ans, test.expected)
		}
	}
}

// run -> go test -v -run="lex"
