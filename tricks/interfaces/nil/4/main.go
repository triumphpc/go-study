package main

import "fmt"

func main() {
	hello(nil) // false, 1
	hello()    // true, 0
}

func hello(a ...interface{}) {
	fmt.Println(a == nil)
	fmt.Println(len(a))
}
