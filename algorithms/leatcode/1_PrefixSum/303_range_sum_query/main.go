// https://leetcode.com/problems/range-sum-query-immutable/
package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	result := 0
	for i := left; i <= right; i++ {
		result += this.nums[i]
	}

	return result
}

// NumArray2 Реализация правильная
type NumArray2 struct {
	nums []int
}

func Constructor2(nums []int) NumArray2 {
	arr := NumArray2{
		nums: make([]int, len(nums)),
	}

	// 1. Сначала подготавливаем суммы всех пар
	arr.nums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr.nums[i] = arr.nums[i-1] + nums[i]
	}

	return arr
}

func (this *NumArray2) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}
	// 2. Вычитаем с правой минус левая
	return this.nums[right] - this.nums[left-1]
}
