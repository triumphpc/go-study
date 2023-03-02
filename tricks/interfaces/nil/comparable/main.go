package main

import "fmt"

func main() {
	var x0 []int
	fmt.Println(x0) // []

	x := []int(nil) // аналог
	fmt.Println(x)  // []

	var y map[int]int
	fmt.Println(y)

	y = map[int]int(nil) // map[]
	fmt.Println(y)       // map[]

	fmt.Println(x[:])   // норм, получает все значения
	fmt.Println(y[344]) // это же map, можно проверять что есть ключ

	z := interface{}(x) // Приведение к интерфейсу типа
	fmt.Println(z)      // []

	//x1 := []int{1, 2, 3}
	//x2 := []int{1, 2, 3}
	//z1 := interface{}(x1) // [1,2,3]
	z2 := interface{}(1)
	//_ = z1 == z1   // нельзя сравнивать panic: runtime error: comparing uncomparable type []int
	fmt.Println(z2 == z2) // ok

	//_ = x1 == x2 // нельзя сравнивать, несравнивамый тип
	//_ = y == y // тоже нельзя

	w := interface{}(y)
	fmt.Println(z == w)  // аа так можно
	fmt.Println(z2 == w) // ok
	//_ = z == z // comparing uncomparable type []int

}
