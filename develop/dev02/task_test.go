package main

import (
	"strconv"
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		hasError bool
	}{
		// success cases
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},

		// error cases
		{"45", "", true},
		{"qwe\\\\5\\", "", true},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			result, err := unpack(tc.input)

			if result != tc.expected {
				t.Errorf("Test case %d: expected \"%s\", got %s", i+1, tc.expected, result)
			}

			if tc.hasError && err == nil || !tc.hasError && err != nil {
				t.Errorf("Test case %d: expected error, got %v", i+1, err)
			}
		})
	}
}
