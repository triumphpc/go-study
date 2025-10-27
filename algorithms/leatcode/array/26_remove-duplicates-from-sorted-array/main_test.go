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
				removeDuplicates(tc.nums)
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
			nums: []int{1, 1, 2},
			val:  3,
		},
		{
			name: "test",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			val:  2,
		},
	}

	for _, tc := range testCases {
		res := removeDuplicates2(tc.nums)
		log.Println(res)
		log.Println(tc.nums)
	}

}

// my
func removeDuplicates(nums []int) int {
	cnt := len(nums)
	for rightIdx := len(nums) - 1; rightIdx > 0; rightIdx-- {
		if nums[rightIdx] == nums[rightIdx-1] {
			cnt--
			for i := rightIdx; i < cnt; i++ {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	return cnt
}

func removeDuplicates2(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0 // points to  the index of last filled position
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i] // O(n) просто проходит по всем и месяет их местами
		}
	}

	return j + 1
}
