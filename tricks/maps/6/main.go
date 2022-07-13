package main

import "fmt"

func main() {
	inc := map[int]bool{1: false, 3: true}

	val, ok := inc[555]

	fmt.Println(val, ok)

}
