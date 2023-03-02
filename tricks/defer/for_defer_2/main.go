package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func(i *int) {
			fmt.Println(*i) // 5 5 5 5
		}(&i)
	}

	for i := 0; i < 3; i++ {
		defer func() { print(i) }() // 3 3 3
	}

	for i := range [3]int{} {
		defer func() { print(i) }() // 2 2 2
	}

	for i := 0; i < 5; i++ {
		defer func(i *int) {
			fmt.Println(*i)
		}(&i)
	}

}
