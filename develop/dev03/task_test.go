package main

import (
	"testing"
)

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
			}
		}
	})
}

func TestCompare(t *testing.T) {
	_ = []struct {
		input    [][]string
		expected [][]string
	}{
		{
			[][]string{{"1"}, {"2"}, {"3"}, {"4"}},
			[][]string{{"1"}, {"2"}, {"3"}, {"4"}},
		},
		{
			[][]string{{"1"}, {"2"}, {"4"}, {"3"}},
			[][]string{{"4"}, {"2K"}, {"3K"}, {"1M"}},
		},
	}

	//t.Run("TestCompare", func(t *testing.T) {
	//
	//	for i := range tc {
	//		result, err := compare(tc[i].input)
	//
	//		for i := range result {
	//			if tc.expected[i][0] != result[i][0] {
	//				t.Errorf("Test case %d: expected \"%s\", got %s", i+1, tc.expected, result)
	//			}
	//		}
	//	}
	//})
}
