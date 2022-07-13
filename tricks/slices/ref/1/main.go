package main

import "fmt"

func main() {
	x := []int{1, 2}
	y := []int{3, 4}
	ref := x

	x = y

	fmt.Printf("%p %v\n", x, x)
	fmt.Printf("%p %v\n", y, y)
	fmt.Printf("%p %v\n", ref, ref)

}
