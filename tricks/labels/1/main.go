package main

import "fmt"

func main() {
	// Выход из цикла по метке
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}

	fmt.Println("out!")
}
