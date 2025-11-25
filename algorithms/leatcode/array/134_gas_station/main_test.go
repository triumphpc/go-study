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
			name:  "merge_sorted",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{3, 4, 5, 1, 2},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(canCompleteCircuit(tc.nums1, tc.nums2))
	}

}

func canCompleteCircuit(gas []int, cost []int) int {
	var currGain int = 0
	var totalGain int = 0
	var answer int = 0

	for i := 0; i < len(gas); i++ {
		totalGain += gas[i] - cost[i]
		currGain += gas[i] - cost[i]
		if currGain < 0 {
			answer = i + 1
			currGain = 0
		}
	}
	if totalGain >= 0 {
		return answer
	} else {
		return -1
	}
}
