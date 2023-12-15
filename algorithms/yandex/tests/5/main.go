package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// 2-SUM
// В данном подходе используем метод двух указателей
// 1. Сначала сортируем
// 2. Выставляем два указателя (лева-право)
// 3. Если сумма эдементов указателей больше числа - смещаем правый влево и наоборот

type TestRun struct {
}

func (*TestRun) Name() string {
	return "2-SUM. Продвинутый 1"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	// 1. Сортируем список
	slices.Sort(in)

	// 2. Установка и обход по указателям
	left := 0
	right := len(in) - 1

	for left < right {
		if in[left]+in[right] == float32(k) { // Если левый указатель и правый равны - то выводим сразу
			result = append(result, in[left], in[right])

			return result, nil
		}

		if in[left]+in[right] > float32(k) { // Если сумма больше нудного - двигаем правую границу
			right--
		} else {
			left++ // // Если сумма больше нудного - двигаем левую
		}
	}

	return result, nil
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))

}

// Вывод индекса а не число
func twoSum(nums []int, target int) []int {
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

// Самое правильное и быстрое решение O(n) O(n) память
// https://www.code-recipe.com/post/two-sum
func twoSumShortAndSimple(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent { // Целевой - текущий = ранее уже сохраненный
			// Ранее сохраненный + текущий = целевой
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
