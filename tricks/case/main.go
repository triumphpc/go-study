package main

import "fmt"

func main() {
	val := 1

	switch val {
	case 1,
		2:
		fmt.Println(val)
		fmt.Println(val)
	default:
		fmt.Println("XXX")

	}

}
