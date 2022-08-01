package main

import "fmt"

func main() {
	//var a []int = nil
	//a, a[0] = []int{1, 9}, 9 //panic: runtime error: index out of range [0] with length 0

	a := []int{123}
	a, a[0] = nil, 9 // []

	fmt.Println(a)

}
