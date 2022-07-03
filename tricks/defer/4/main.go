package main

import "fmt"

func main() {

	fmt.Println(test())  // 1
	fmt.Println(test2()) // 2
	fmt.Println(test3()) // 2
	fmt.Println(test4()) // 1

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

func test3() (x int) {
	// в случае func() значение считывается исходя их выполнение всей фукнции. Т.е. в defer будет значение x, которое на выходе (1)
	defer func() {
		fmt.Println("here 2", x) // 1
		x++
	}()

	fmt.Println("here 1")
	x = 1

	return x
}

func test4() (x int) {
	// в данном случае берется значение на момент запуска функции = 0. Не возвращает значение, поэтому  = 1
	defer func() {
		x++
	}()

	x = 2

	return 0
}
