package main

import "fmt"

func main() {
	type A = int
	var a A
	a = 10
	fmt.Printf("%T", a)
}
