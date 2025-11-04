package main

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
			nums1: []int{2, 1, 2, 1, 0, 1, 2},
			k:     3,
		},
	}

	for _, tc := range testCases {
		res := maxProfit(tc.nums1)
		log.Println(res)
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
				_ = maxProfit(tc.nums1)
			}
		})
	}

}

// Benchmark
// Benchmark/test
// Benchmark/test-14         	305904724	         3.830 ns/op
// PASS

// Benchmark
// Benchmark/test
// Benchmark/test-14         	691933282	         1.656 ns/op
// PASS
func maxProfit(prices []int) int {
	minV := -1
	profit := 0
	for i := 0; i < len(prices); i++ {
		curr := prices[i]
		s := curr - minV
		if minV == -1 || curr < minV {
			minV = curr
		} else if s > profit {
			profit = s
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	minprice := int(^uint(0) >> 1) // инициализация максимальным int
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
