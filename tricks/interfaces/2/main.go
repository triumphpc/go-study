package main

import "fmt"

type S struct {
	a, b int
}

// Указатели в интерфейсах не равны
func main() {
	s1 := &S{1, 2}
	s2 := &S{1, 2}
	x := interface{}(s1)
	y := interface{}(s2)

	fmt.Printf("Value of the tank interface is: %v %T %p \n", x, x, x)
	fmt.Printf("Value of the tank interface is: %v %T %p \n", y, y, y)

	// But not equal because its point
	fmt.Println(x == y)     // false
	fmt.Println(s1 == s2)   // false
	fmt.Println(*s1 == *s2) // true

}
