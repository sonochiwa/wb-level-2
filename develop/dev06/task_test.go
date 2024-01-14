package main

import (
	"reflect"
	"testing"
)

func TestSelectedFields(t *testing.T) {
	test := struct {
		name     string
		input    [][]string
		fields   []int
		expected [][]string
	}{
		name:     "test selected fields",
		input:    [][]string{{"1", "3"}, {"2", "2"}, {"3", "1"}},
		fields:   []int{2},
		expected: [][]string{{"3"}, {"2"}, {"1"}},
	}

	t.Run(test.name, func(t *testing.T) {
		result := selectedFields(test.input, test.fields)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	})
}
