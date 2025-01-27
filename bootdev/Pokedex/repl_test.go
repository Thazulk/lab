package main

import "testing"

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
			input:    "Luke, I am your father",
			expected: []string{"luke,", "i", "am", "your", "father"},
		},
		{
			input:    "offci is da king",
			expected: []string{"offci", "is", "da", "king"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input '%s': expected %d words, got %d words",
				c.input, len(c.expected), len(actual))
			continue
		}
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Expected %s but got %s", expectedWord, word)
			}
			// and fail the test

		}
	}

}
