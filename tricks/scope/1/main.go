package main

import "fmt"

func main() {
	i, s := 2, []int{}

	for i = range s {
	}
	fmt.Println(i, s) // 2 []

	for i = 0; i < len(s); i++ {
	}
	fmt.Println(i, s) // 0 []

	s = append(s, 1, 2, 3, 4)
	for i = range s {
	}

	fmt.Println(i, s) // 3 [1 2 3 4]

	for i = 0; i < len(s); i++ {
	}
	fmt.Println(i, s) // 4 [1 2 3 4]

}
