package main

import "fmt"

const X = 2

func main() {
	const (
		X = X + X // тут переопределение в области видимости фунакции X = 4
		Y         // тут уже принцип iota X + X (4 + 4)
		Z
		F
	)

	fmt.Println(X, Y, Z, F)

	const L = 2

	const (
		K = L
		K1

		K2
	)

	fmt.Println(K, K1, K2) // 2 2 2

}
