// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

package benchmarking

import (
	"log"
	"sort"
	"testing"
)

func BenchmarkMerge(b *testing.B) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
	}{
		{
			name:  "merge_sorted",
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				mergeMy(tc.nums1, 3, tc.nums2, 3)
			}
		})
	}

}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
	}{
		{
			name:  "merge_sorted",
			nums1: []int{7, 5, 3, 0, 0, 0},
			nums2: []int{1, 2, 5},
		},
	}

	for _, tc := range testCases {
		mergeMy(tc.nums1, 3, tc.nums2, 3)
	}

}

func mergeMy(nums1 []int, m int, nums2 []int, n int) {
	// 1. Сначала просто вставим второй слайс в первый
	for j := 0; j < n; j++ {
		nums1[m+j] = nums2[j]
	}

	// 2. Просто сортировка без выделение памяти
	sort.Ints(nums1)

	log.Println(nums1)
}
