package _38_product_of_array_except_self

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{nums: []int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"t1", args{nums: []int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkProductExceptSelf
// BenchmarkProductExceptSelf/sequential_numbers
// BenchmarkProductExceptSelf/sequential_numbers-14         	     813	   1237816 ns/op
// BenchmarkProductExceptSelf/random_numbers_(no_zeros)
// BenchmarkProductExceptSelf/random_numbers_(no_zeros)-14  	     948	   1253321 ns/op
// BenchmarkProductExceptSelf/all_ones
// BenchmarkProductExceptSelf/all_ones-14                   	     960	   1225528 ns/op
// BenchmarkProductExceptSelf/repeating_pattern
// BenchmarkProductExceptSelf/repeating_pattern-14          	     975	   1224591 ns/op
// BenchmarkProductExceptSelf/small_input
// BenchmarkProductExceptSelf/small_input-14                	75929144	        14.48 ns/op
// PASS
func BenchmarkProductExceptSelf(b *testing.B) {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Размер тестовых данных
	size := 1000

	// Создаем тестовые данные разных типов
	nums1 := make([]int, size) // последовательные числа
	for i := range nums1 {
		nums1[i] = i + 1
	}

	nums2 := make([]int, size) // случайные числа (избегаем нулей, чтобы не было деления на 0)
	for i := range nums2 {
		nums2[i] = rand.Intn(10) + 1 // числа от 1 до 10
	}

	nums3 := make([]int, size) // все единицы (крайний случай)
	for i := range nums3 {
		nums3[i] = 1
	}

	nums4 := make([]int, size) // повторяющийся паттерн
	pattern := []int{1, 2, 3, 4, 5}
	for i := range nums4 {
		nums4[i] = pattern[i%len(pattern)]
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("sequential numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf(nums1)
		}
	})

	b.Run("random numbers (no zeros)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf(nums2)
		}
	})

	b.Run("all ones", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf(nums3)
		}
	})

	b.Run("repeating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf(nums4)
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{1, 2, 3, 4}
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf(smallNums)
		}
	})
}

// BenchmarkProductExceptSelf2
// BenchmarkProductExceptSelf2/sequential_numbers
// BenchmarkProductExceptSelf2/sequential_numbers-14         	  226554	      4894 ns/op
// BenchmarkProductExceptSelf2/random_numbers_(no_zeros)
// BenchmarkProductExceptSelf2/random_numbers_(no_zeros)-14  	  224798	      5150 ns/op
// BenchmarkProductExceptSelf2/all_ones
// BenchmarkProductExceptSelf2/all_ones-14                   	  224815	      5193 ns/op
// BenchmarkProductExceptSelf2/repeating_pattern
// BenchmarkProductExceptSelf2/repeating_pattern-14          	  244954	      4981 ns/op
// BenchmarkProductExceptSelf2/small_input
// BenchmarkProductExceptSelf2/small_input-14                	40897413	        28.91 ns/op
// PASS
func BenchmarkProductExceptSelf2(b *testing.B) {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Размер тестовых данных
	size := 1000

	// Создаем тестовые данные разных типов
	nums1 := make([]int, size) // последовательные числа
	for i := range nums1 {
		nums1[i] = i + 1
	}

	nums2 := make([]int, size) // случайные числа (избегаем нулей, чтобы не было деления на 0)
	for i := range nums2 {
		nums2[i] = rand.Intn(10) + 1 // числа от 1 до 10
	}

	nums3 := make([]int, size) // все единицы (крайний случай)
	for i := range nums3 {
		nums3[i] = 1
	}

	nums4 := make([]int, size) // повторяющийся паттерн
	pattern := []int{1, 2, 3, 4, 5}
	for i := range nums4 {
		nums4[i] = pattern[i%len(pattern)]
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("sequential numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf2(nums1)
		}
	})

	b.Run("random numbers (no zeros)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf2(nums2)
		}
	})

	b.Run("all ones", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf2(nums3)
		}
	})

	b.Run("repeating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf2(nums4)
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{1, 2, 3, 4}
		for i := 0; i < b.N; i++ {
			_ = productExceptSelf2(smallNums)
		}
	})
}
