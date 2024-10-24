package main

import "testing"

func TestCleanInput(t *testing.T) {
	test := []struct {
		input  string
		output []string
	}{
		{
			input:  "Hello",
			output: []string{"hello"},
		},
		{
			input:  "Help",
			output: []string{"help"},
		},
		{
			input:  "EXIT",
			output: []string{"exit"},
		},
		{
			input: "Hello WOrld",
			output: []string{
				"hello",
				"world",
			},
		},
	}
	for _, tc := range test {
		result := cleanInput(tc.input)
		actual := len(result)
		if expectedLen := len(tc.output); actual != expectedLen {
			t.Errorf("Lengths dont match: %v vs %v", actual, expectedLen)
			continue
		}

		for i := range actual {

			if result[i] != tc.output[i] {
				t.Errorf("cleanInput(%v)= %v, expected  %v", tc.input, result, tc.output)
			}
		}
	}
}
