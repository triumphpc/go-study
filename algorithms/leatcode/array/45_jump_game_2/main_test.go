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
			nums1: []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			k:     3,
		},
	}

	for _, tc := range testCases {

		log.Println(canJump(tc.nums1))
	}

}

func Benchmark(b *testing.B) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
	}{
		{
			name:  "test",
			nums1: []int{7, 5, 3, 5, 0, 0, 0},
			k:     3,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				canJump(tc.nums1)
			}
		})
	}

}

// вот лучшее решение
func canJump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var maxIdx, i, step int
	for i < len(nums)-1 {
		nextIdx, tmpIdx := 0, 0
		for j := 0; j < nums[i]; j++ {
			tmpIdx = i + j + 1
			if tmpIdx < len(nums) && tmpIdx+nums[tmpIdx] > maxIdx {
				maxIdx = tmpIdx + nums[tmpIdx]
				nextIdx = tmpIdx
			}
			if tmpIdx == len(nums)-1 {
				nextIdx = tmpIdx
			}
		}

		i = nextIdx
		step++
	}

	return step
}
