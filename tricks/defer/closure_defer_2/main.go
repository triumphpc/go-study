package main

import "fmt"

func main() {
	fmt.Println(test()) // 1

}

func test() (result int) {
	defer func() {
		result++
		fmt.Println(result) // 2
	}()

	return 1
}
