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
			nums1: []int{29, 51, 87, 87, 72, 12},
			nums2: []int{3, 4, 5, 1, 2},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(candy(tc.nums1))
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
				_ = candy(tc.nums1)
			}
		})
	}

}

// 67374384	        17.85 ns/op
// мое решение побило 70%
func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	candies[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else if ratings[i] < ratings[i-1] {
			candies[i] = 1
			for j := i - 1; j >= 0; j-- {
				if candies[j] > candies[j+1] {
					break
				}

				if ratings[j] == ratings[j+1] {
					break
				}

				if candies[j] == candies[j+1] {
					candies[j] = candies[j] + 1
				}
			}
		} else {
			candies[i] = 1
		}

	}

	var res int
	for i := 0; i < len(candies); i++ {
		res += candies[i]
	}

	return res
}
