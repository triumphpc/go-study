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

	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			goto loopS
		}
	}
loopS:

	fmt.Println("out!")
}
