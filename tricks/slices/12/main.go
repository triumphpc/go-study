package main

import "fmt"

// в какой строке будет паника

func main() {
	x := []int(nil)
	fmt.Println(x, len(x), cap(x)) // [] 0 0

	x = []int{}
	fmt.Println(x, len(x), cap(x)) // [] 0 0

	_ = x[:]

	y := map[int]int(nil)
	fmt.Println(y) // map[]

	z := interface{}(x)
	fmt.Println(z) // []

	w := interface{}(y)
	fmt.Println(w) // map[]

	_ = z == w // false
	fmt.Printf("here: %T -  %v\n", w, w)
	fmt.Printf("here: %T -  %v", z, z)

	_ = z == z // panic: runtime error: comparing uncomparable type []int

}
