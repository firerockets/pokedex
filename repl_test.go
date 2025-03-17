package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected size %d. Found %d", len(c.expected), len(actual))
			t.Fail()
		}

		for i, word := range actual {
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected word %s. Found %s", expectedWord, word)
				t.Fail()
			}
		}
	}
}
