package main

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)

	tests := []struct {
		left, right int
		want        int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
	}

	for _, tt := range tests {
		got := na.SumRange(tt.left, tt.right)
		if got != tt.want {
			t.Errorf("SumRange(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
		}
	}
}

// BenchmarkSumRange
// BenchmarkSumRange/small_range
// BenchmarkSumRange/small_range-14         	46212038	        24.04 ns/op
// BenchmarkSumRange/medium_range
// BenchmarkSumRange/medium_range-14        	  527389	      2265 ns/op
// BenchmarkSumRange/large_range
// BenchmarkSumRange/large_range-14         	   51907	     23756 ns/op
// PASS
func BenchmarkSumRange(b *testing.B) {
	// Создаем большой массив для тестирования
	size := 100000
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i % 100 // заполняем случайными значениями
	}

	na := Constructor(nums)

	b.ResetTimer() // Сбрасываем таймер, чтобы не учитывать время подготовки

	// Тестируем разные диапазоны
	b.Run("small range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(100, 200)
		}
	})

	b.Run("medium range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(100, 10000)
		}
	})

	b.Run("large range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(0, size-1)
		}
	})
}

func TestNumArray2(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor2(nums)

	tests := []struct {
		left, right int
		want        int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
	}

	for _, tt := range tests {
		got := na.SumRange(tt.left, tt.right)
		if got != tt.want {
			t.Errorf("SumRange(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
		}
	}
}

// BenchmarkSumRange2
// BenchmarkSumRange2/small_range
// BenchmarkSumRange2/small_range-14         	1000000000	         0.2427 ns/op
// BenchmarkSumRange2/medium_range
// BenchmarkSumRange2/medium_range-14        	1000000000	         0.2274 ns/op
// BenchmarkSumRange2/large_range
// BenchmarkSumRange2/large_range-14         	1000000000	         0.2267 ns/op
// PASS
func BenchmarkSumRange2(b *testing.B) {
	// Создаем большой массив для тестирования
	size := 100000
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i % 100 // заполняем случайными значениями
	}

	na := Constructor2(nums)

	b.ResetTimer() // Сбрасываем таймер, чтобы не учитывать время подготовки

	// Тестируем разные диапазоны
	b.Run("small range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(100, 200)
		}
	})

	b.Run("medium range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(100, 10000)
		}
	})

	b.Run("large range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = na.SumRange(0, size-1)
		}
	})
}
