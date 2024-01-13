package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	test := struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		name:  "test find anagrams",
		input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
		expected: map[string][]string{
			"пятак":  {"пятак", "пятка", "тяпка"},
			"листок": {"листок", "слиток", "столик"},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		result := findAnagrams(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %s: expected \"%s\", got \"%s\"", test.name, test.expected, result)
			return
		}
	})
}

func TestSortWord(t *testing.T) {
	test := struct {
		name     string
		input    string
		expected string
	}{
		name:     "test sort words",
		input:    "пятак",
		expected: "акптя",
	}

	t.Run(test.name, func(t *testing.T) {
		result := sortWord(test.input)
		if test.expected != result {
			t.Errorf("Test case %s: expected \"%s\", got \"%s\"", test.name, test.expected, result)
			return
		}
	})
}
