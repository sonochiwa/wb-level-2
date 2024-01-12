package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// Test sort algorithm
//func TestSortStrings(t *testing.T) {
//
//}

// Test -k flag
//func TestSetColumn(t *testing.T) {
//
//}

// Test -n flag
//func TestSortNumbers(t *testing.T) {
//
//}

// Test -r flag
//func TestReverseData(t *testing.T) {
//
//}

// Test -u flag
func TestDedupeData(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name:     "no duplicates",
			input:    [][]string{{"data1"}, {"data2"}, {"data3"}},
			expected: [][]string{{"data1"}, {"data2"}, {"data3"}},
		},
		{
			name:     "with duplicates",
			input:    [][]string{{"duplicate"}, {"duplicate"}},
			expected: [][]string{{"duplicate"}},
		},
		{
			name:     "multiple duplicates",
			input:    [][]string{{"multi"}, {"multi"}, {"multi"}},
			expected: [][]string{{"multi"}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := dedupeData(tc.input)

			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("TestDedupeData(%s) input \"%v\", exptected \"%v\"", tc.name, input, tc.expected)
			}

		})
	}
}

// Test -M flag
func TestSortMonths(t *testing.T) {
	test := struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		name: "test sort months",
		input: [][]string{
			{"february"},
			{"april"},
			{"january"},
			{"march"},
		},
		expected: [][]string{
			{"january"},
			{"february"},
			{"march"},
			{"april"},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		result := sortMonths(test.input)

		for i := range result {
			if result[i][0] != test.expected[i][0] {
				t.Errorf("Test case %d: expected \"%s\", got \"%s\"", i+1, test.expected, result)
				return
			}
		}
	})
}

// Test -b flag
//func TestTrimSpace(t *testing.T) {
//
//}

// Test -c flag
func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		input1   [][]string
		input2   [][]string
		expected error
	}{
		{
			name:     "test case without error",
			input1:   [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			input2:   [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			expected: nil,
		},
		{
			name:     "test case with error",
			input1:   [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			input2:   [][]string{{"b"}, {"a"}, {"c"}, {"d"}},
			expected: errors.New(fmt.Sprint("disorder:a")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			err := compare(tc.input1, tc.input2)
			if err != nil {
				if err.Error() != tc.expected.Error() {
					t.Errorf("Test case %s: expected \"%s\", got \"%s\"", tc.name, tc.expected, err)
					return
				}
			}
		})
	}
}

// Test -H flag
func TestHumanNumericSort(t *testing.T) {
	test := struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		name:     "test success case",
		input:    [][]string{{"1M"}, {"4"}, {"3K"}, {"2K"}},
		expected: [][]string{{"4"}, {"2K"}, {"3K"}, {"1M"}},
	}

	t.Run(test.name, func(t *testing.T) {
		result := humanNumericSort(test.input)

		for i := range result {
			if test.expected[i][0] != result[i][0] {
				t.Errorf("Test case %s: expected \"%s\", got %s", test.name, test.expected, result)
				return
			}
		}
	})
}
