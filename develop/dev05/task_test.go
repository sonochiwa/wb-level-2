package main

import (
	"reflect"
	"testing"
)

func TestGetIndexes(t *testing.T) {
	test := struct {
		name     string
		input    []string
		pattern  string
		expected map[int]bool
	}{
		name:     "test get indexes",
		input:    []string{"fake", "hello world", "real fake"},
		pattern:  "hello world",
		expected: map[int]bool{0: false, 1: true, 2: false},
	}

	t.Run(test.name, func(t *testing.T) {
		result := getIndexes(test.pattern, test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %s: expected \"%v\", got \"%v\"", test.name, test.expected, result)
			return
		}
	})
}

func TestAfter(t *testing.T) {

}

func TestBefore(t *testing.T) {

}

func TestContext(t *testing.T) {

}

func TestInvert(t *testing.T) {

}

func TestCountStrings(t *testing.T) {

}
