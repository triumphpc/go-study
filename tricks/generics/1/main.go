package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(max(1, 5))

	var ux, uy uint
	ux = 4
	uy = 8

	fmt.Println(maxGeneric(ux, uy))

	fmt.Println(maxGeneric2(ux, uy))

}

// Обычная функция
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Работа с обобщенным типом (дженериком)
// [T any] T - обобщенный тип, any - ограничение
// comparable - констрейн для сравнения равны ли они друг другу (не позволит сдлеать больше или меньше)
// Для больше/меньше нужно свой контрейнт определять

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func maxGeneric[T comparable](x, y T) T {
	if x == y {
		return x
	}

	return y
}

func maxGeneric2[T Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// Можно импортировать стандартные контрейнты constraints.Ordered
func maxGeneric3[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}
