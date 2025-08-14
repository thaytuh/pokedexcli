package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  charmander  Bulbasaur  PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "THIs   is   just   a   test HAHAHA",
			expected: []string{"this", "is", "just", "a", "test", "hahaha"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual slice length does not match expected slice length")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word does not match expected word")
			}
		}
	}
}