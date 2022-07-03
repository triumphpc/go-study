package main

import "fmt"

type S struct{}

func initPointer() *S {
	var s *S

	return s
}

func initType() S {
	var s S

	return s
}

func initInterface() interface{} {
	var s S
	return s

}

func main() {

	fmt.Println(initPointer() == nil) // true

	//fmt.Println(initType() == nil) // error

	fmt.Println(initInterface() == nil) // false

	var i interface{}

	if i == nil {
		fmt.Println("nil") // true
	}
}
