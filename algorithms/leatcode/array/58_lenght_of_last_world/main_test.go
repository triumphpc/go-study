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
			name: "a",
			k:    3,
		},
	}

	for _, tc := range testCases {
		log.Println(lengthOfLastWord(tc.name))
	}

}

func lengthOfLastWord(s string) int {

	var res int
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]

		if ch == 32 && res == 0 {
			continue
		}

		if ch == 32 {
			break
		}

		res++
	}

	return res

}
