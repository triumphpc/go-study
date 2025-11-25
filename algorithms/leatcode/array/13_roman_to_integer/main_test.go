package _69_majority_element

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
	}{
		{
			name:  "MCMXCIV",
			nums1: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			nums2: []int{3, 4, 5, 1, 2},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(romanToInt(tc.name))
	}

}

func romanToInt(s string) int {
	var value = map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	var sum int
	var last rune

	for i, r := range s {
		v, l := value[r], value[last]
		sum += v
		last = r
		if i > 0 && v > l {
			sum -= l * 2
		}
	}

	return sum
}
