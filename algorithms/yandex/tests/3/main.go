package test

// Метод скользящего среднего
// Оптимизация алгоритма

type TestRun struct {
}

func (*TestRun) Name() string {
	return "Метод скользящего среднего. Оптимизация"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {
	maxCap := cap(in)
	var current float32

	for _, v := range in[:k] {
		current += v
	}
	result = append(result, current/float32(k))

	for idx := 0; idx < maxCap-k; idx++ {
		current -= in[idx]
		current += in[idx+k]

		result = append(result, current/float32(k))
	}

	return result, nil
}
