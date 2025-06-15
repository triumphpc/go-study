package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{nums: []int{1, 1, 1}, k: 3}, 3},
		{"t1", args{nums: []int{1, 1, 1, 0, 0, 1}, k: 3}, 5},
		{"t1", args{nums: []int{1, 1, 3, 2, 11, 1, 1, 1, 1, 1}, k: 15}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum3(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubarraySum1
// BenchmarkSubarraySum1/sequential_numbers
// BenchmarkSubarraySum1/sequential_numbers-14         	    6339	    377547 ns/op
// BenchmarkSubarraySum1/random_numbers
// BenchmarkSubarraySum1/random_numbers-14             	    6080	    379309 ns/op
// BenchmarkSubarraySum1/all_zeros
// BenchmarkSubarraySum1/all_zeros-14                  	    6297	    191340 ns/op
// BenchmarkSubarraySum1/repeating_pattern
// BenchmarkSubarraySum1/repeating_pattern-14          	    2922	    376618 ns/op
// BenchmarkSubarraySum1/small_input
// BenchmarkSubarraySum1/small_input-14                	73678582	        18.30 ns/op
// PASS
func BenchmarkSubarraySum1(b *testing.B) {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем тестовые данные разных типов
	size := 1000               // Уменьшено по сравнению с 100000, так как алгоритм O(n^3) очень медленный
	nums1 := make([]int, size) // последовательные числа
	for i := range nums1 {
		nums1[i] = i + 1
	}

	nums2 := make([]int, size) // случайные числа
	for i := range nums2 {
		nums2[i] = rand.Intn(200) - 100 // числа от -100 до 100
	}

	nums3 := make([]int, size) // все нули
	for i := range nums3 {
		nums3[i] = 0
	}

	nums4 := make([]int, size) // повторяющийся паттерн
	pattern := []int{1, -1, 2, -2, 3, -3}
	for i := range nums4 {
		nums4[i] = pattern[i%len(pattern)]
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("sequential numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum(nums1, 42) // произвольный k
		}
	})

	b.Run("random numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum(nums2, 0) // сумма 0
		}
	})

	b.Run("all zeros", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum(nums3, 0) // сумма 0
		}
	})

	b.Run("repeating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum(nums4, 1) // сумма 1
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{1, 1, 1, 2, 2, -3, 1, 1, 1}
		for i := 0; i < b.N; i++ {
			_ = subarraySum(smallNums, 3) // несколько подмассивов с суммой 3
		}
	})
}

// BenchmarkSubarraySum2
// BenchmarkSubarraySum2/sequential_numbers
// BenchmarkSubarraySum2/sequential_numbers-14         	      25	  42229413 ns/op
// BenchmarkSubarraySum2/random_numbers
// BenchmarkSubarraySum2/random_numbers-14             	      27	  42160151 ns/op
// BenchmarkSubarraySum2/all_zeros
// BenchmarkSubarraySum2/all_zeros-14                  	      28	  40938354 ns/op
// BenchmarkSubarraySum2/repeating_pattern
// BenchmarkSubarraySum2/repeating_pattern-14          	      27	  41350046 ns/op
// BenchmarkSubarraySum2/small_input
// BenchmarkSubarraySum2/small_input-14                	19345386	        61.29 ns/op
// PASS
func BenchmarkSubarraySum2(b *testing.B) {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем тестовые данные разных типов
	size := 1000               // Уменьшено по сравнению с 100000, так как алгоритм O(n^3) очень медленный
	nums1 := make([]int, size) // последовательные числа
	for i := range nums1 {
		nums1[i] = i + 1
	}

	nums2 := make([]int, size) // случайные числа
	for i := range nums2 {
		nums2[i] = rand.Intn(200) - 100 // числа от -100 до 100
	}

	nums3 := make([]int, size) // все нули
	for i := range nums3 {
		nums3[i] = 0
	}

	nums4 := make([]int, size) // повторяющийся паттерн
	pattern := []int{1, -1, 2, -2, 3, -3}
	for i := range nums4 {
		nums4[i] = pattern[i%len(pattern)]
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("sequential numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum2(nums1, 42) // произвольный k
		}
	})

	b.Run("random numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum2(nums2, 0) // сумма 0
		}
	})

	b.Run("all zeros", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum2(nums3, 0) // сумма 0
		}
	})

	b.Run("repeating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum2(nums4, 1) // сумма 1
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{1, 1, 1, 2, 2, -3, 1, 1, 1}
		for i := 0; i < b.N; i++ {
			_ = subarraySum2(smallNums, 3) // несколько подмассивов с суммой 3
		}
	})
}

// это лучший
// BenchmarkSubarraySum3
// BenchmarkSubarraySum3/sequential_numbers
// BenchmarkSubarraySum3/sequential_numbers-14         	  293248	      4013 ns/op
// BenchmarkSubarraySum3/random_numbers
// BenchmarkSubarraySum3/random_numbers-14             	  297478	      4067 ns/op
// BenchmarkSubarraySum3/all_zeros
// BenchmarkSubarraySum3/all_zeros-14                  	    8289	    144841 ns/op
// BenchmarkSubarraySum3/repeating_pattern
// BenchmarkSubarraySum3/repeating_pattern-14          	  292888	      4048 ns/op
// BenchmarkSubarraySum3/small_input
// BenchmarkSubarraySum3/small_input-14                	74453810	        18.98 ns/op
// PASS
func BenchmarkSubarraySum3(b *testing.B) {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем тестовые данные разных типов
	size := 1000               // Уменьшено по сравнению с 100000, так как алгоритм O(n^3) очень медленный
	nums1 := make([]int, size) // последовательные числа
	for i := range nums1 {
		nums1[i] = i + 1
	}

	nums2 := make([]int, size) // случайные числа
	for i := range nums2 {
		nums2[i] = rand.Intn(200) - 100 // числа от -100 до 100
	}

	nums3 := make([]int, size) // все нули
	for i := range nums3 {
		nums3[i] = 0
	}

	nums4 := make([]int, size) // повторяющийся паттерн
	pattern := []int{1, -1, 2, -2, 3, -3}
	for i := range nums4 {
		nums4[i] = pattern[i%len(pattern)]
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("sequential numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum3(nums1, 42) // произвольный k
		}
	})

	b.Run("random numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum3(nums2, 0) // сумма 0
		}
	})

	b.Run("all zeros", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum3(nums3, 0) // сумма 0
		}
	})

	b.Run("repeating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = subarraySum3(nums4, 1) // сумма 1
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{1, 1, 1, 2, 2, -3, 1, 1, 1}
		for i := 0; i < b.N; i++ {
			_ = subarraySum3(smallNums, 3) // несколько подмассивов с суммой 3
		}
	})
}
