package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

func Double[T Number](value T) T {
	return value * 2
}

func DotProduct[T Number](s1, s2 []T) T {
	if len(s1) != len(s2) {
		panic("DotProduct: slices of unequal length")
	}
	var r T
	for i := range s1 {
		r += s1[i] * s2[i]
	}
	return r
}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Invoke Double
	fmt.Println(Double(23))
	fmt.Println(Double(23.23))
	fmt.Println(Double(-2323.3434))

	// Invoke DotProduct
	i := []int{1, 2, 3}
	j := []int{4, 5, 6}
	fmt.Println(DotProduct(i, j))

	// Invoke Sum
	ints := map[string]int64{
		"first":   23,
		"second":  565,
		"third":   755,
		"fourth":  766,
		"fifth":   8977,
		"sixth":   70433,
		"seventh": 4339222,
	}
	fmt.Println(Sum(ints))

}
