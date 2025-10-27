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
	}{
		{
			name:  "merge_sorted",
			nums1: []int{7, 5, 3, 5, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		res := majorityElementMy(tc.nums1)
		log.Println(res)
	}

}

func Benchmark(b *testing.B) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
	}{
		{
			name:  "test",
			nums1: []int{7, 5, 3, 5, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				majorityElementMy(tc.nums1)
			}
		})
	}

}

// Benchmark
// Benchmark/test
// Benchmark/test-14         	19513820	        60.38 ns/op
// PASS
func majorityElementMy(nums []int) int {
	k := make(map[int]int, len(nums))
	var r, m int
	for i := 0; i < len(nums); i++ {
		k[nums[i]]++
		if k[nums[i]] > r {
			r = k[nums[i]]
			m = nums[i]

		}
	}

	return m
}

// Benchmark
// Benchmark/test
// Benchmark/test-14         	15325756	        77.92 ns/op
// PASS
func majorityElement2(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		if _, ok := counts[num]; ok {
			counts[num]++
		} else {
			counts[num] = 1
		}
	}
	for num, count := range counts {
		if count > len(nums)/2 {
			return num
		}
	}
	return 0
}

// Benchmark/test
// Benchmark/test-14         	24205362	        49.56 ns/op
// PASS
func majorityElementBest(nums []int) int {
	res := 0
	max := -1
	freq := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		freq[val]++

		if cnt, _ := freq[val]; cnt > max {
			max = cnt
			res = val
		}
	}

	return res
}
