package test

import "golang.org/x/exp/slices"

// Бинарный поиск
// Задача найти элемент в массиве
// Сложность O(log n)
// 1. Определяем границы, по которым делаем поиск min/max
// 2. Сортируем
// 3. Берем середину, если равно - сразу отдаем
// 4. Если средний элемент меньше - беррем правую часть для следующей итерации и наоборот

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

		if in[mid] == float32(k) { // Если середина равна нужному - выдаем ее сразу
			result = append(result, float32(k))

			return
		}

		if in[mid] < float32(k) { // Если значение меньше нужного - берем правую часть для следующей итерации
			low = mid + 1
		} else {
			high = mid - 1 // или левую
		}
	}

	return result, nil
}
