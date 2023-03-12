package test

import "golang.org/x/exp/slices"

// Линейный поиск
// Задача найти элемент в массиве

type TestRun struct {
}

func (*TestRun) Name() string {
	return "Бинарный поиск"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	// 1. Определяем границы по которым делаем поиск
	low := 0
	high := len(in) - 1

	// 2. Сортируем список
	slices.Sort(in)

	for low <= high { // Пока не сократится до одного элемента
		mid := (low + high) / 2

		if in[mid] == float32(k) {
			result = append(result, float32(k))

			return
		}

		if in[mid] < float32(k) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result, nil
}
