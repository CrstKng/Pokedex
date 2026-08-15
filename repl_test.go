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
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "JohnY dOEY",
			expected: []string{"johny", "doey"},
		},
		{
			input: "I  need  more  codE ",
			expected: []string{"i", "need", "more", "code"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("not the same number words: actual: %d, expected: %d", len(actual), len(c.expected))
			continue
		}
		for i, actualWord := range actual {
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("words %s and %s don't match", actualWord, expectedWord)
				t.Fail()
			}
		}
	}
}
