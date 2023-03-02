package main

import "fmt"

func main() {
	// Скомпилируется ли код?
	//x := []int{123}
	//fmt.Println(len(x), cap(x)) // 1 1
	//x = nil
	////fmt.Println(len(x), cap(x)) // 0 0
	//x, x[0] = nil, 456
	//
	////x[0] = 455 x = nil
	//fmt.Println(len(x), cap(x)) // 0 0
	//x, x[0] = []int{555}, 456   // тут будет panic, потому что x стал nil

	// А как с массивами
	var x *[3]int
	//x = &[...]int{0: 1} // В таком формате нужно обязательно указать последний элемент
	x = &[...]int{2: 3} // В таком формате нужно обязательно указать последний элемент, важен размер

	x[0] = 0
	fmt.Println(len(x), cap(x)) // 3 3
	x = nil

	fmt.Println(len(x), cap(x)) // 3 3

	x = &[...]int{2: 3}
	x[0] = 0
	fmt.Println(len(x), cap(x))
	x = nil // указатель на массив, ничего не меняет!!!

	fmt.Println(len(x), cap(x))

	var y [3]int
	y = [...]int{2: 3}
	fmt.Println(len(y), cap(y)) // 3 3
	//y = nil                     // error, cant use

	fmt.Println(len(y), cap(y)) // 3 3

}
