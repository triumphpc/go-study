package main

// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.
//
//
//
//Example 1:
//
//Input: nums = [0,1]
//Output: 2
//Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
//Example 2:
//
//Input: nums = [0,1,0]
//Output: 2
//Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
//Example 3:
//
//Input: nums = [0,1,1,1,1,1,0,0,0]
//Output: 6
//Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.
//
//
//Constraints:
//
//1 <= nums.length <= 105
//nums[i] is either 0 or 1.

func findMaxLength1(nums []int) int {
	// Инициализируйте переменную count для отслеживания разности между количеством 1 и 0, и
	//переменную max_length для хранения максимальной длины подмассива. Создайте хеш-таблицу map
	//для хранения первых встреч каждого значения count. Добавьте начальное значение (0, -1) в хеш-таблицу.
	countMap := map[int]int{0: -1}
	count, maxLength := 0, 0

	// Итеративно пройдите по массиву nums. На каждой итерации обновляйте значение count (увеличивайте на 1 для 1 и
	//уменьшайте на 1 для 0). Если текущее значение count уже существует в хеш-таблице, вычислите длину подмассива между текущим индексом и
	//индексом из хеш-таблицы. Обновите max_length, если текущий подмассив длиннее.
	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		if prevIndex, exists := countMap[count]; exists {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			// Если текущее значение count не существует в хеш-таблице, добавьте его с текущим индексом. После завершения итерации верните max_length.
			countMap[count] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
