package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}

}
func main() {
	f := F(5)
	fmt.Println(f()) // 6

	defer func() {
		fmt.Println(f()) // 10
	}()

	defer func() {
		fmt.Println(f()) // 9 (берется по возврату)
	}()

	defer fmt.Println(f()) // 7 (берется в момент инициализации)

	fmt.Println(f()) // 8

}
