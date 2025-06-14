package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_findMaxLength1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{nums: []int{0, 1, 0}}, 2},
		{"t2", args{nums: []int{0, 1, 0, 1, 1, 1}}, 4},
		{"t3", args{nums: []int{0, 0, 1, 1, 1, 1, 1, 0, 0, 0}}, 10},
		{"t3", args{nums: []int{0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}, 8},
		{"t3", args{nums: []int{0, 1, 1, 1, 1, 1, 0, 0, 0}}, 6},
		{"t3", args{nums: []int{0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength1(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindMaxLength1
// BenchmarkFindMaxLength1/alternating_pattern
// BenchmarkFindMaxLength1/alternating_pattern-14         	   48240	     22702 ns/op
// BenchmarkFindMaxLength1/random_pattern
// BenchmarkFindMaxLength1/random_pattern-14              	   53046	     22741 ns/op
// BenchmarkFindMaxLength1/long_sequences
// BenchmarkFindMaxLength1/long_sequences-14              	   51151	     22539 ns/op
// BenchmarkFindMaxLength1/small_input
// BenchmarkFindMaxLength1/small_input-14                 	449083657	         2.621 ns/op
// PASS
func BenchmarkFindMaxLength1(b *testing.B) {
	// Создаем большие массивы для тестирования разных сценариев
	size := 100000
	nums1 := make([]int, size) // чередующиеся 0 и 1
	for i := range nums1 {
		nums1[i] = i % 2
	}

	nums2 := make([]int, size) // случайные 0 и 1
	rand.Seed(time.Now().UnixNano())
	for i := range nums2 {
		nums2[i] = rand.Intn(2)
	}

	nums3 := make([]int, size) // длинные последовательности
	for i := range nums3 {
		if i < size/2 {
			nums3[i] = 0
		} else {
			nums3[i] = 1
		}
	}

	b.ResetTimer()

	// Тестируем разные сценарии
	b.Run("alternating pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findMaxLength1(nums1)
		}
	})

	b.Run("random pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findMaxLength1(nums2)
		}
	})

	b.Run("long sequences", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findMaxLength1(nums3)
		}
	})

	b.Run("small input", func(b *testing.B) {
		smallNums := []int{0, 1, 0, 1, 1, 1, 0, 0, 1, 1}
		for i := 0; i < b.N; i++ {
			_ = findMaxLength1(smallNums)
		}
	})
}
