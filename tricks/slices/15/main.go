package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%v %T %v %v\n", a, a, cap(a), len(a)) // slice

	b := [...]int{1, 2, 3}
	fmt.Printf("%v %T %v %v\n", b, b, cap(b), len(b)) // array

	c := b
	b[0] = 42

	fmt.Printf("%v %T %v %v\n", b, b, cap(b), len(b)) // array
	fmt.Printf("%v %T %v %v\n", c, c, cap(c), len(c)) // array

	fmt.Println(c, b)

}
