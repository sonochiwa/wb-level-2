package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		result   string
		expected string
		error    error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", nil},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result, err := unpack(tc.result)
			if err != nil {
				fmt.Println(err)
			}
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
	//result, err := unpack("a4bc2d5e")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//expected := "aaaabccddddde"
	//
	//if result != expected {
	//	t.Errorf("Expected %s, but got %s", expected, result)
	//}
}
