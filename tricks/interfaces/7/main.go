package main

import (
	"fmt"
	"log"
)

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
	log.Printf("%T %v\n", c, c) //  main.C {0 0}

	_, bf := c.(interface{ f() })      // true  приведение типа
	_, bg := c.(interface{ g() })      // false приведение типа (так как их несколько у наследников)
	bm, ok := c.(interface{ m() int }) // true приведение типа

	log.Printf("%T %v\n", bm, bm) //  main.C {0 0}

	fmt.Println(bf, bg, ok, bm.m()) // true false true 9

}
