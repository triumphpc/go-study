package main

import "fmt"

func main() {
	var x0 []int
	fmt.Println(x0) // []

	x := []int(nil) // аналог
	y := map[int]int(nil)
	_ = x[:]
	_ = y[3]
	z := interface{}(x) // []
	//x1 := []int{1, 2, 3}
	//z1 := interface{}(x1) // [1,2,3]
	//_ = z1 == z1 // нельзя сравнивать
	//_ = x1 == x1 // нельзя сравнивать
	//_ = y == y // тоже нельзя

	w := interface{}(y)
	_ = z == w // аа так можно
	_ = z == z // comparing uncomparable type []int

}
