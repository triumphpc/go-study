package test

// Метод скользящего среднего
//С его помощью мы снизим шумы в данных и сгладим график, выделив общую тенденцию изменения числа запросов.
// Напишем самый простой псевдокод функции, реализующей метод скользящего среднего.
//Самое очевидное и простое решение какой-либо задачи называется «наивный алгоритм» — это устоявшийся термин.
// Наивный алгоритм

type TestRun struct {
}

func (*TestRun) Name() string {
	return "Метод скользящего среднего"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	maxCap := cap(in)
	for idx := 0; idx < maxCap-k; idx++ {
		endIdx := idx + k
		if endIdx > maxCap {
			endIdx = maxCap - 1
		}
		// Просматриваем окно ширинок k
		var current float32
		for _, v := range in[idx:endIdx] {
			current += v
		}

		result = append(result, current/float32(k))
	}

	return result, nil
}
