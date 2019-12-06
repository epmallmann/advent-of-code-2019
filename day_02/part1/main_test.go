package main

import "testing"

func TestIntcode(t *testing.T) {
	var testValues = []struct {
		input  string
		output string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50"},
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, testVal := range testValues {
		output := calculateIntcode(testVal.input)

		if output != testVal.output {
			t.Errorf("error calculating intcode. got [%s], want [%s]", output, testVal.output)
		}
	}
}
