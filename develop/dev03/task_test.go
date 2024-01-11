package main

import (
	"errors"
	"fmt"
	"testing"
)

// Test sort algorithm
func TestSortStrings(t *testing.T) {

}

// Test -k flag
func TestSetColumn(t *testing.T) {

}

// Test -n flag
func TestSortNumbers(t *testing.T) {

}

// Test -r flag
func TestReverseData(t *testing.T) {

}

// Test -u flag
func TestDedupeData(t *testing.T) {

}

// Test -M flag
func TestSortMonths(t *testing.T) {
	tc := struct {
		input    [][]string
		expected [][]string
	}{
		[][]string{
			{"february"},
			{"april"},
			{"january"},
			{"march"},
		},
		[][]string{
			{"january"},
			{"february"},
			{"march"},
			{"april"},
		},
	}

	result := sortMonths(tc.input)

	for i := range result {
		if result[i][0] != tc.expected[i][0] {
			t.Errorf("Test case %d: expected \"%s\", got \"%s\"", i+1, tc.expected, result)
			return
		}
	}
}

// Test -b flag
func TestTrimSpace(t *testing.T) {

}

// Test -c flag
func TestCompare(t *testing.T) {
	tc := []struct {
		input1   [][]string
		input2   [][]string
		expected error
	}{
		{
			[][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			[][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			nil,
		},
		{
			[][]string{{"a"}, {"b"}, {"c"}, {"d"}},
			[][]string{{"b"}, {"a"}, {"c"}, {"d"}},
			errors.New(fmt.Sprint("disorder:a")),
		},
	}

	t.Run("TestCompare", func(t *testing.T) {
		for i := range tc {
			err := compare(tc[i].input1, tc[i].input2)
			if err != nil {
				if err.Error() != tc[i].expected.Error() {
					t.Errorf("Test case %d: expected \"%s\", got \"%s\"", i+1, tc[i].expected, err)
					return
				}
			}
		}
	})
}

// Test -H flag
func TestHumanNumericSort(t *testing.T) {
	tc := struct {
		input    [][]string
		expected [][]string
	}{
		[][]string{{"1M"}, {"4"}, {"3K"}, {"2K"}},
		[][]string{{"4"}, {"2K"}, {"3K"}, {"1M"}},
	}

	t.Run("TestHumanNumericSort", func(t *testing.T) {
		result := humanNumericSort(tc.input)

		for i := range result {
			if tc.expected[i][0] != result[i][0] {
				t.Errorf("Test case %d: expected \"%s\", got %s", i+1, tc.expected, result)
				return
			}
		}
	})
}
