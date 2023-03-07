package test

// 2-SUM
// Дан массив целых чисел numbers и целое число X; нужно найти в массиве два элемента, сумма которых равняется X.
// Гарантируется, что такие элементы в массиве есть.
// Наивный алгоритм. Перебор по одному каждый.
// x x x x x x x x x
//
//	x x x x x x x x
//	  x x x x x x x
// Сколько попыток? N^2/2

type TestRun struct {
}

func (*TestRun) Name() string {
	return "2-SUM. Наивный"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if (in[i] + in[j]) == float32(k) {
				result = append(result, in[i], in[j])

				return result, nil
			}

		}
	}

	return result, nil
}
