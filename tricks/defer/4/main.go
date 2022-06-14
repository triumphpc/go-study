package main

import "fmt"

func main() {

	fmt.Println(test())  // 1
	fmt.Println(test2()) // 2

}

func test() int {
	x := 1
	defer func() {
		fmt.Println("here 2", x)
		x++
	}()

	fmt.Println("here 1")

	return x
}

func test2() (x int) {
	x = 1

	defer func() {
		fmt.Println("here 2", x) // 1
		x++
	}()

	x++
	fmt.Println("here 1")

	return x
}
