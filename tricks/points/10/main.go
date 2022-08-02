package main

import "fmt"

type MyType int

func main() {
	tt := MyType(1)

	var tt2 *MyType

	tt3 := MyType(55)
	tt2 = &tt3

	fmt.Println(tt)

	fmt.Println(tt2)

	test(tt)
	test(*tt2)

}

func test(tt MyType) {
	fmt.Println(tt)

}
