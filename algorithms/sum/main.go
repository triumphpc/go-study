package main

import (
	"fmt"
	"reflect"
)

func twoSum(input []int, target int) []int {
	// write solution here
	// тут просто решение написать, но в самом деле я ее не давал никому
	return nil
}

func main() {
	testCases := []struct {
		input    []int
		target   int
		expected []int
	}{
		{
			input:    []int{2, 7, 11, 15},
			target:   18,
			expected: []int{1, 2},
		},
		{
			input:    []int{0, 0, 0, 4, 0, 0, 7, 0, 0, 0},
			target:   11,
			expected: []int{3, 6},
		},
		{
			input:    []int{3, 2, 4},
			target:   7,
			expected: []int{0, 2},
		},
		{
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, testCase := range testCases {
		if !reflect.DeepEqual(twoSum(testCase.input, testCase.target), testCase.expected) {
			fmt.Printf("ERROR: case %v\n", testCase)
			return
		}
	}
	fmt.Println("Done!")
}
