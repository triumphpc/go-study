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
			nums1: []int{1, 5, 3, 2, 1, 0},
			k:     3,
		},
	}

	for _, tc := range testCases {
		res := maxProfitMy(tc.nums1)
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
				_ = maxProfitMy(tc.nums1)
			}
		})
	}

}

func maxProfitMy(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		curr := prices[i]
		prev := prices[i-1]

		if curr > prev {
			profit += curr - prev
		}
	}

	return profit
}

// хуже решение
func maxProfit2(prices []int) int {
	i := 0
	valley := prices[0]
	peak := prices[0]
	maxprofit := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxprofit += peak - valley
	}
	return maxprofit
}
