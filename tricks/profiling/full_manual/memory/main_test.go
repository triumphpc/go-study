// https://leetcode.com/problems/two-sum/description/
// https://blog.stackademic.com/profiling-go-applications-in-the-right-way-with-examples-e784526e9481

package benchmarking

import (
	"testing"

	"golang.org/x/exp/slices"
)

func BenchmarkTwoSum(b *testing.B) {
	// Тестовые случаи разного размера
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "small_slice",
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			name:   "medium_slice",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
		},
		{
			name:   "large_slice",
			nums:   make([]int, 1000),
			target: 1998,
		},
	}

	// Заполняем большой слайс
	for i := range testCases[2].nums {
		testCases[2].nums[i] = i + 1
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				twoSum(tc.nums, tc.target)
			}
		})
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	// Return an empty slice if no solution is found
	return []int{}
}

func BenchmarkTwoSumHashTable(b *testing.B) {
	// Тестовые случаи разного размера
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "small_slice",
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			name:   "medium_slice",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
		},
		{
			name:   "large_slice",
			nums:   make([]int, 1000),
			target: 1998,
		},
	}

	// Заполняем большой слайс
	for i := range testCases[2].nums {
		testCases[2].nums[i] = i + 1
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				twoSumHashTable(tc.nums, tc.target)
			}
		})
	}
}

func twoSumHashTable(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	// Return an empty slice if no solution is found
	return []int{}
}

func BenchmarkTwoSumOnePassHashTable(b *testing.B) {
	// Тестовые случаи разного размера
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "small_slice",
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			name:   "medium_slice",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
		},
		{
			name:   "large_slice",
			nums:   make([]int, 1000),
			target: 1998,
		},
	}

	// Заполняем большой слайс
	for i := range testCases[2].nums {
		testCases[2].nums[i] = i + 1
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				twoSumOnePassHashTable(tc.nums, tc.target)
			}
		})
	}
}

func twoSumOnePassHashTable(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	// Return an empty slice if no solution is found
	return []int{}
}

func BenchmarkTwoMy(b *testing.B) {
	// Тестовые случаи разного размера
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "small_slice",
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			name:   "medium_slice",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
		},
		{
			name:   "large_slice",
			nums:   make([]int, 1000),
			target: 1998,
		},
	}

	// Заполняем большой слайс
	for i := range testCases[2].nums {
		testCases[2].nums[i] = i + 1
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Сброс таймера перед каждым тестом
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				twoSumMy(tc.nums, tc.target)
			}
		})
	}
}

func twoSumMy(nums []int, target int) []int {
	idxs := make(map[int][]int, len(nums))
	for i := range nums {
		if idxs[nums[i]] == nil {
			idxs[nums[i]] = make([]int, 0, 1)
		}
		idxs[nums[i]] = append(idxs[nums[i]], i)
	}

	slices.Sort(nums)

	// 2. Установка и обход по указателям
	left := 0
	right := len(nums) - 1
	result := make([]int, 0, 2)

	for left < right {
		if nums[left]+nums[right] == target { // Если левый указатель и правый равны - то выводим сразу
			leftIdx := idxs[nums[left]][:1][0]
			idxs[nums[left]] = idxs[nums[left]][1:]
			rightIdx := idxs[nums[right]][:1][0]
			idxs[nums[right]] = idxs[nums[right]][1:]

			result = append(result, leftIdx, rightIdx)

			return result
		}

		if nums[left]+nums[right] > target { // Если сумма больше нудного - двигаем правую границу
			right--
		} else {
			left++ // // Если сумма больше нудного - двигаем левую
		}
	}

	return result
}
