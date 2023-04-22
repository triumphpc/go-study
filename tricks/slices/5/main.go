package main

import "fmt"

func fn(sl []int) {
	fmt.Printf("(2)%v - %p\n", sl, sl) // (2)[1 2 3 4] - 0xc00004c020
	sl[1] = 555
}

func fn2(sl []int) {
	fmt.Printf("(2)%v - %p\n", sl, sl) // (2)[1 555 3 4 666] - 0xc000020140
	sl[1] = 666
}

func main() {
	sl := []int{1, 2, 3, 4}
	fmt.Printf("(1)%v - %p\n", sl, sl) // (1)[1 2 3 4] - 0xc00004c020
	fn(sl)

	fmt.Printf("(3)%v - %p\n", sl, sl) // (3)[1 555 3 4] - 0xc00004c020

	sl = append(sl, 666)
	fmt.Printf("(4)%v - %p\n", sl, sl) // (4)[1 555 3 4 666] - 0xc00001a100

	fn2(sl)
	fmt.Printf("(5)%v - %p\n", sl, sl) // (5)[1 666 3 4 666] - 0xc000020140

}
