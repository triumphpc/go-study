package main

import "fmt"

// Вывод:
//2
//3 33
//1 22

func main() {
	defer func() {
		fmt.Println("1", recover())
	}()

	defer func() {
		fmt.Println("2")
		defer fmt.Println("3", recover())
		panic(22)
	}()

	defer recover()

	panic(33)
}
