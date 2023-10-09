package test

// 2-SUM
// Карта хранения веса. Самый быстрый вариант.
// 1. Отнимаем каждый элемент по очереди из числа
// 2. Если уже есть такой в карте, то возвращем его
// 3. Если нет - устанавливаем

type TestRun struct {
}

func (*TestRun) Name() string {
	return "2-SUM. Продвинутый 2. С вспомогательной картой"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	set := make(map[float32]bool, len(in))

	for idx := range in {
		y := float32(k) - in[idx]

		// Если уже есть в карте, то возвращаем
		if set[y] {
			result = append(result, in[idx], y)

			return
		}

		set[in[idx]] = true
	}

	return result, nil
}
