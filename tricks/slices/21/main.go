package main

import "fmt"

func main() {
	var x *[9]int
	x = &[...]int{8: 1, 0: 2}
	fmt.Println(x, len(x), cap(x)) // &[2 0 0 0 0 0 0 0 1] 9 9

	y := x[1:3]                    // [0 0]
	fmt.Println(y, cap(y), len(y)) //[0 0] 8 2

	z := y[5:7]                    // [0 0]
	fmt.Println(z, cap(z), len(z)) //[0 0] 3 2
	// С базового берет
}
