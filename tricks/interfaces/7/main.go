package main

import "fmt"

var o = fmt.Print

type A int

func (A) g()     {}
func (A) m() int { return 1 }

type B int

func (B) g() {}
func (B) f() {}

type C struct {
	A
	B
}

func (C) m() int {
	return 9
}

func main() {

	var c interface{} = C{}
	_, bf := c.(interface{ f() })  // true  приведение типа
	_, bg := c.(interface{ g() })  // false приведение типа
	bm := c.(interface{ m() int }) // false приведение типа

	fmt.Println(bf, bg, bm.m()) // true false 9

}
