// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

package benchmarking

import (
	"log"
	"testing"
)

func Benchmark(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
		val  int
	}{
		{
			name: "test",
			nums: []int{3, 2, 2, 3},
			val:  2,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				removeElement(tc.nums, tc.val)
			}
		})
	}

}

func Benchmark2(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
		val  int
	}{
		{
			name: "test",
			nums: []int{3, 2, 2, 3},
			val:  2,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				removeElement2(tc.nums, tc.val)
			}
		})
	}

}

func Test(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		val  int
	}{
		{
			name: "test",
			nums: []int{3, 2, 2, 3},
			val:  3,
		},
		{
			name: "test",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
		},
	}

	for _, tc := range testCases {
		res := removeElement(tc.nums, tc.val)
		log.Println(res)
		log.Println(tc.nums)
	}

}

// my
func removeElement(nums []int, val int) int {
	start := 0
	final := len(nums)

	for {
		if start == final {
			break
		}

		if nums[start] == val {
			nums[start] = 0
			nums[start], nums[final-1] = nums[final-1], nums[start]
			final--
		} else {
			start++
		}
	}

	return start
}

// Best
func removeElement2(nums []int, val int) int {
	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	return writeIndex
}
