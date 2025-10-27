package _0_remove_dublicates_2

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
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
			val:  9,
		},
		{
			name: "test",
			nums: []int{1, 1, 1, 2, 2, 3},
			val:  5,
		},
		{
			name: "test",
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			val:  7,
		},
	}

	for _, tc := range testCases {
		res := removeDuplicates(tc.nums)
		require.Equal(t, res, tc.val)
		log.Println(res, tc.nums)
	}

}

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if j < 2 || nums[i] > nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
