package main

import (
	"fmt"
)

func main() {
	// Инициализация мапы с интовыми значениями
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Инициализация мапы с float значениями
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

}

// SumIntsOrFloats суммирует значения мамы m.
// Поддерживает int64 и float64 типы элементов в мапе.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s

}
