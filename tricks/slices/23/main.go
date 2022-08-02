package main

import "fmt"

func main() {
	var x *[9]int
	x = &[...]int{8: 0, 0: 2}      // 8 элемент со значением 0
	fmt.Println(x, cap(x), len(x)) //&[0 0 0 0 0 0 0 0 0] 9 9

	y := x[1:3]                    // с 1(включительно) по 3 (не включительно)
	fmt.Println(y, cap(y), len(y)) //[0 0] 8 2

	z := y[5:7]
	fmt.Println(z, cap(z), len(z)) //[0 0] 3 2

}
