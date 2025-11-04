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
			nums1: []int{2, 0},
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

// рабочий вариант я его не понял
//type Index int
//
//const (
//	GOOD    Index = 1
//	BAD     Index = 2
//	UNKNOWN Index = 0
//)
//
//var memo []Index
//
//func canJumpFromPosition(position int, nums []int) bool {
//	if memo[position] != UNKNOWN {
//		return memo[position] == GOOD
//	}
//	furthestJump := position + nums[position]
//	if furthestJump > len(nums)-1 {
//		furthestJump = len(nums) - 1
//	}
//	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
//		if canJumpFromPosition(nextPosition, nums) {
//			memo[position] = GOOD
//			return true
//		}
//	}
//	memo[position] = BAD
//	return false
//}
//
//func canJump(nums []int) bool {
//	memo = make([]Index, len(nums))
//	for i := range memo {
//		memo[i] = UNKNOWN
//	}
//	memo[len(nums)-1] = GOOD
//	return canJumpFromPosition(0, nums)
//}

// вот лучшее решение
func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}
