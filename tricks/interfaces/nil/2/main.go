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

type IF interface {
	F()
}

func (s S) F() {

}

func initInterfacePointer() interface{} {
	var s *S

	return s
}

func initInterfaceType() IF {
	var s S

	return s
}

func initInterfaceTypePointer() IF {
	var s *S

	return s
}

func main() {
	fmt.Println(initInterfacePointer() == nil) // false

	fmt.Println(initInterfaceType() == nil) // false

	fmt.Println(initInterfaceTypePointer() == nil) // false

	fmt.Println(initPointer() == nil) // true

	fmt.Println(initType()) // {}

	//fmt.Println(initType() == nil) // error, тип структуры нельзя сравниться с nil

	fmt.Println(initInterface() == nil) // false

	var i interface{}

	if i == nil {
		fmt.Println("nil") // true
	}
}
