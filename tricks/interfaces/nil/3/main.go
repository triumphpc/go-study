package main

import "fmt"

type MyType struct {
	string
}

func main() {
	var evt *MyType

	evt = &MyType{"fdfd"}

	fmt.Println(evt == nil)
}
