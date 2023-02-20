package main

import "fmt"

// Что выведет?
func main() {
	var i int
	for i = 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
			break // 00 , 10
		}
	}

	fmt.Println(i) // 2
}
