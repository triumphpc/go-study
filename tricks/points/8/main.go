package main

import "fmt"

// в функцию передается по значению
// создается копия
func add(arr []int, x int) {
	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 0 0 0 0 0 0 0 0 0] 0xc000010060 10 10
	// видно, что это новый слайс в другой области памяти, но ссылается на один базовый массив

	arr[1] = 1 // до аллокации ссылаются на общий базовый массив

	arr = append(arr, x) // тут добавляется новый элемент, выделяется новая область памяти

	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 1 0 0 0 0 0 0 0 0 10] 0xc000010060 11 20

	arr[9] = 9
	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 1 0 0 0 0 0 0 0 9 10] 0xc000010060 11 20

	arr[1] = 11
	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 11 0 0 0 0 0 0 0 9 10] 0xc000010060 11 20

}

func main() {
	arr := make([]int, 10)
	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 0 0 0 0 0 0 0 0 0] 0xc000010030 10 10

	add(arr, 10)

	fmt.Printf("%+v %p %d %d \n", arr, &arr, len(arr), cap(arr)) // [0 1 0 0 0 0 0 0 0 0] 0xc000010030 10 10

}
