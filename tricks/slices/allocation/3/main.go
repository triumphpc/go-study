package main

import "fmt"

func main() {
	arr := make([]int, 1, 1)

	fmt.Printf("%v %p\n", arr, &arr) // [] 0xc0000a4018
	add(arr, 100)                    // (1) Передается в функцию по значению

	fmt.Printf("%v %p\n", arr, &arr) // [123] 0xc00000c030 (3 видим, что поменялось только значение базового массива)

	fmt.Printf("%v %v\n", cap(arr), len(arr)) // 1 1 длинна и вместимость не меняются

}

func add(n []int, v int) {
	fmt.Printf("%v %p\n", n, &n) // [] 0xc0000a4048 (ссылается на другую область памяти для слайса, так как передано по значению)
	n[0] = 123                   // (2) Меняется значение в базовом массиве

	fmt.Printf("%v %v\n", cap(n), len(n)) // 1 1
	n = append(n, v)                      // при добавлении элемента меняется капасити слайса, переданного в функцию
	// Так же создается новые базовый массив, на который ссылается данный слайс, но при этом первое значение
	// ссылается на изначальный базовый массив

	fmt.Printf("%v %p\n", n, &n)          // [123 100] 0xc00000c060
	fmt.Printf("%v %v\n", cap(n), len(n)) // 2 2 (видим, что меняется капасити и значения слайса)
}