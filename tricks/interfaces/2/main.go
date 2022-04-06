package main

import "fmt"

type S struct {
	a, b int
}

func main() {
	// Interfaces type S
	x := interface{}(&S{1, 2})
	y := interface{}(&S{1, 2})

	fmt.Printf("Value of the tank interface is: %v %T %p \n", x, x, x)
	fmt.Printf("Value of the tank interface is: %v %T %p \n", y, y, y)

	// But not equal because its point
	fmt.Println(x == y) // false

}
