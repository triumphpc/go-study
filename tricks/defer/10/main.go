package main

import (
	"fmt"
	"strconv"
)

func F(n int) func() int {
	return func() int {
		fmt.Println("It's:" + strconv.Itoa(n))
		n++

		return n
	}
}

func main() {
	f := F(5)
	f() // 6

	fmt.Println(f()) // 7
	fmt.Println(f()) // 8

	defer func() {
		fmt.Println("defer")
		fmt.Println(f()) // 11
	}()

	defer fmt.Println(f()) // вызывает функцию но не выводит занчение
	// становится 9
	// после выводит уже ее

	fmt.Println("here")
	f() // 10

}
