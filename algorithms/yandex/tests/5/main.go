package test

import (
	"golang.org/x/exp/slices"
)

// 2-SUM
// В данном подходе используем метод двух указателей

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
		currentSum := in[left] + in[right]
		if in[left]+in[right] == float32(k) {
			result = append(result, in[left], in[right])

			return result, nil
		}

		if currentSum > float32(k) {
			right--
		} else {
			left++
		}
	}

	return result, nil
}
