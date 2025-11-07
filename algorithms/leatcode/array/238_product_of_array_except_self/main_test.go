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
			nums1: []int{-1, 1, 0, -3, 3},
			k:     3,
		},

		{
			name:  "merge_sorted",
			nums1: []int{1, 2, 3, 4},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(productExceptSelf(tc.nums1))
	}

}

// долго по времени не выполнятся
func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	for i := 0; i < l; i++ {
		res[i] = 1
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}

			if nums[j] == 0 {
				res[i] = 0
				break
			}

			res[i] = res[i] * nums[j]
		}
	}

	return res
}

// Рабочий вариант
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	L := make([]int, length)
	R := make([]int, length)
	answer := make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}

	return answer
}

// лучшее решение но тут с делением, не верно
func productExceptSelf3(nums []int) []int {
	n := len(nums)
	product := nums[0]
	zeroCount := 0
	for i := 1; i < n; i++ {
		if nums[i] == 0 && zeroCount == 0 {
			zeroCount++
			continue
		}
		if product == 0 && zeroCount == 0 {
			product = nums[i]
			zeroCount++
			continue
		}
		product *= nums[i]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			res[i] = product
		} else if zeroCount == 0 {
			res[i] = product / nums[i]
		}
	}
	return res
}
