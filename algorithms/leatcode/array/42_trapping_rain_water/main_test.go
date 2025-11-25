package _69_majority_element

import (
	"log"
	"math"
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
			nums1: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			nums2: []int{3, 4, 5, 1, 2},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(trap2(tc.nums1))
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
				_ = trap(tc.nums1)
			}
		})
	}

}

// не верно работает
func trap(height []int) int {
	level := 0
	curr := 0
	res := 0
	now := 0

	for i := 0; i < len(height); i++ {
		curr = height[i]

		if level > 0 {
			if curr >= level {
				level = curr
				res += int(math.Abs(float64(now)))
				now = 0
			} else {
				now = now + (curr - level)
			}

		} else {
			level = curr
		}
	}

	return res

}

func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans := 0
	size := len(height)
	left_max, right_max := make([]int, size), make([]int, size)
	left_max[0] = height[0]
	for i := 1; i < size; i++ {
		if height[i] > left_max[i-1] {
			left_max[i] = height[i]
		} else {
			left_max[i] = left_max[i-1]
		}
	}
	right_max[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		if height[i] > right_max[i+1] {
			right_max[i] = height[i]
		} else {
			right_max[i] = right_max[i+1]
		}
	}
	for i := 1; i < size-1; i++ {
		if left_max[i] < right_max[i] {
			ans += left_max[i] - height[i]
		} else {
			ans += right_max[i] - height[i]
		}
	}
	return ans
}
