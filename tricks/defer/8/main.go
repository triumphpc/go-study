package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func(i *int) {
			fmt.Println(*i) // 5 5 5 5 5
		}(&i)
	}
}
