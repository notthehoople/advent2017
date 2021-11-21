package main

import (
	"testing"
)

func TestCalcSteps(t *testing.T) {
	var testsA = []struct {
		input    int
		expected int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	// Test for Part A
	for _, test := range testsA {
		if output := calcSteps(test.input, false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}