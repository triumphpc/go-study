package main

import "fmt"

func main() {
	x := []int{1, 3}
	y := []int{2, 4}

	ref := x
	x = y

	fmt.Printf("%v %v %v", x, y, ref)

}
