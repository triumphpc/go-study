// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150
package _38_product_of_array_except_self

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	var start bool

	for i := 0; i < len(nums); i++ {
		start = false
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if !start {
				result[i] = nums[j]
				start = true

				continue
			}

			result[i] = result[i] * nums[j]
		}
	}

	return result
}

// 1⃣Инициализация массивов L и R: Инициализируйте два пустых массива L и R. Массив L будет содержать произведение всех
// чисел слева от i, а массив R будет содержать произведение всех чисел справа от i. Заполните массив L, установив L[0]
// равным 1, а для остальных элементов используйте формулу L[i] = L[i-1] * nums[i-1]. Заполните массив R, установив R[length-1]
// равным 1, а для остальных элементов используйте формулу R[i] = R[i+1] * nums[i+1].
//
// 2⃣Заполнение массивов L и R: Пройдите два цикла для заполнения массивов L и R. В первом цикле заполните массив L,
// начиная с L[0] и двигаясь вправо. Во втором цикле заполните массив R, начиная с R[length-1] и двигаясь влево.
//
// 3⃣Формирование результата: Пройдите по исходному массиву и для каждого элемента i вычислите произведение всех элементов,
// кроме nums[i], используя L[i] * R[i]. Сохраните результат в массиве answer и верните его.
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
