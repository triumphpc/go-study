package test

// Линейный поиск
// Задача найти элемент в массиве
// Сложность O(n) - линейная зависимость
type TestRun struct {
}

func (*TestRun) Name() string {
	return "Линейный поиск"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	// Это перебор по одному элементу

	for _, v := range in {
		if float32(k) == v {
			result = append(result, v)

			return
		}
	}

	return result, nil
}
