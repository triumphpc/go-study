package main

import "fmt"

func main() {
	var list []int
	fmt.Println(list == nil)    // true
	fmt.Println(len(list) == 0) // true

	list = []int{}
	println(list == nil)        // false
	fmt.Println(len(list) == 0) // true

}
