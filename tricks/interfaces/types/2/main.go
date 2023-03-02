package main

import (
	"fmt"
	"reflect"
)

func main() {
	type A = int
	var a A

	fmt.Printf("%s\n", reflect.ValueOf(a)) // %!s(int=0)
	fmt.Printf("%v\n", reflect.TypeOf(a))  // int

	a = 10
	fmt.Printf("%s\n", reflect.ValueOf(a)) // %!s(int=0)
	fmt.Printf("%v\n", reflect.TypeOf(a))  //int
}
