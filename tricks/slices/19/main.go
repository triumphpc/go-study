package main

import "fmt"

func main() {
	//var a []int = nil (len 0 cap 0)
	//a, a[0] = []int{1, 9}, 9 //panic: runtime error: index out of range [0] with length 0
	// Справа на лево проставляется

	a := []int{123}
	a, a[0] = nil, 9 // []
	// Тут норм

	fmt.Println(a)

}
