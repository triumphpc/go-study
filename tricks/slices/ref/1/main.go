package main

import "fmt"

func main() {
	x := []int{1, 2}
	y := []int{3, 4}
	ref := x

	fmt.Printf("%p %v\n", x, x)
	fmt.Printf("%p %v\n", y, y)
	fmt.Printf("%p %v\n", ref, ref)

	x[0] = 55

	x = y // x начал ссылаться на y память

	fmt.Printf("%p %v\n", x, x)
	fmt.Printf("%p %v\n", y, y)
	fmt.Printf("%p %v\n", ref, ref)

}
