package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%v %T %v %v\n", a, a, cap(a), len(a)) // [1 2 3] []int 3 3

	b := [...]int{1, 2, 3}
	fmt.Printf("%v %T %v %v\n", b, b, cap(b), len(b)) // [1 2 3] [3]int 3 3

	c := b // копия
	b[0] = 42

	fmt.Printf("%v %T %v %v\n", b, b, cap(b), len(b)) // [42 2 3] [3]int 3 3
	fmt.Printf("%v %T %v %v\n", c, c, cap(c), len(c)) // [1 2 3] [3]int 3 3

	z := a // ссылка на один базовый
	z[0] = 42
	fmt.Printf("%v %T %v %v\n", z, z, cap(z), len(z)) // [42 2 3] []int 3 3
	fmt.Printf("%v %T %v %v\n", a, a, cap(a), len(a)) // [42 2 3] []int 3 3

}
