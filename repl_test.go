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
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HELLO WORLD ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " heLLo WorlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		words := cleanInput(c.input)
		if len(words) != len(c.expected) {
			t.Errorf("Unequal length between '%v' & '%v'", words, c.expected)
			continue
		}
		for i, word := range words {
			if word != c.expected[i] {
				t.Errorf("'%v' != '%v'", word, c.expected[i])
			}
		}
	}

}
