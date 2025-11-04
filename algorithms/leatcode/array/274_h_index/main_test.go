// Определение простыми словами:
// У исследователя h-индекс = N, если:
//
// # У него есть N статей, на каждую из которых сослались как минимум N раз
//
// # И при этом остальные статьи имеют не более N цитирований каждая
//
// Как считать на примерах:
// Пример 1: citations = [3,0,6,1,5]
// Сортируем по убыванию: [6,5,3,1,0]
//
// Сравниваем позицию и количество цитирований:
//
// 1-я статья: 6 цитирований ≥ 1? ✅
//
// 2-я статья: 5 цитирований ≥ 2? ✅
//
// 3-я статья: 3 цитирования ≥ 3? ✅
//
// 4-я статья: 1 цитирование ≥ 4? ❌
//
// Находим максимальное h: h = 3, потому что есть 3 статьи с ≥3 цитирований
package _69_majority_element

import (
	"log"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
	}{
		{
			name:  "merge_sorted",
			nums1: []int{1, 3, 1},
			k:     3,
		},
	}

	for _, tc := range testCases {
		log.Println(hIndex(tc.nums1))
	}

}

func Benchmark(b *testing.B) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
	}{
		{
			name:  "test",
			nums1: []int{7, 5, 3, 5, 0, 0, 0},
			k:     3,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				hIndex(tc.nums1)
			}
		})
	}

}

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	index := 0
	for i := 0; i < len(citations); i++ {
		idx := citations[i]

		if idx >= i+1 {
			index++
		}
	}

	return index
}
