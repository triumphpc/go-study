package main

import "fmt"

// Можно делить на 0
func main() {
	f := 500.0
	fmt.Printf("float: %v\n", f/0) // float: +Inf (бесконечность)

	i := 5
	fmt.Printf("int: %v\n", i/0)
}
