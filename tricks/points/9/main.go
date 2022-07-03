package main

import "fmt"

// в функцию передается по значению
// создается копия
func add(arr []int, x int) {
	fmt.Printf("%v %p\n", arr, &arr)
	arr = append(arr, x)
	fmt.Printf("%v %p\n", arr, &arr)
}

func main() {
	arr := make([]int, 10)
	fmt.Printf("%v %p\n", arr, &arr)
	add(arr, 10)
	fmt.Printf("%v %p\n", arr, &arr)

}
