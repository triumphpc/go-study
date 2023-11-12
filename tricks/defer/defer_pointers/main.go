package main

import (
	"fmt"
)

func main() {
	var i1 int = 10
	var k = 20
	var i2 = &k

	defer printInt("test1", i1)        // 10
	defer printInt("test2", *i2)       // 20
	defer printIntPointer("test3", i2) // 2020
	i1 = 30
	defer printInt("test4", i1) // 30

	i1 = 1010
	*i2 = 2020

}

func printInt(v string, i int) {
	fmt.Println(v, i)

}

func printIntPointer(v string, i *int) {
	fmt.Println(v, *i)
	ff := map[int]int{}

}
