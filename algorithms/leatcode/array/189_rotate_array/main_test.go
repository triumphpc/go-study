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
			nums1: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
		},
	}

	for _, tc := range testCases {
		rotate2(tc.nums1, tc.k)
		log.Println(tc.nums1)
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
				rotate(tc.nums1, tc.k)
			}
		})
	}

}

func Benchmark2(b *testing.B) {
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
				rotate2(tc.nums1, tc.k)
			}
		})
	}

}

// my умер по времени
func rotate(nums []int, k int) {
	for k > 0 {
		for i := len(nums) - 1; i > 0; i-- {
			nums[i], nums[i-1] = nums[i+-1], nums[i]
		}
		k--
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		l := (i + k) % n // запоминаем позицию с остатком от деления
		a[l] = nums[i]
	}
	for i := 0; i < n; i++ {
		nums[i] = a[i]
	}
}
