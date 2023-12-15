package main

import (
	"fmt"
	"unsafe"
)

const (
	_ = iota
	Avito
	There
	One
	Mmmm = iota
	Kkk
)

const (
	Test  = "test" // test
	Test2 = 2      // 2

	state1 = iota + 1 // 3
	state2            // 4
	state3
)

const (
	fixeState = iota + 1 // 1
	fixStat2             //2
)

const (
	a, b = iota, iota + 1     // 0 1
	c, d = iota + 2, iota + 3 // 2 3
)

func main() {
	fmt.Println(Avito, There, One, Mmmm, Kkk) // 1 2 3 4 5

	testIOTA3()
}

func testIOTA1() {
	const (
		p1 = iota // какой тип переменной будет?
		q1 = iota
		r1 = iota
	)
	fmt.Printf("%T %T %T", p1, q1, r1) // int int int
}

func testIOTA2() {
	const (
		p2, q2, r2 = iota, iota, iota
	)

	fmt.Println(p2, q2, r2) // 0 0 0
}

func testIOTA3() {
	const (
		p3 byte = iota
		q3      // какой тип переменной будет? какое будет значение?
		r3
	)

	fmt.Printf("%T %T %T %v", p3, q3, r3, unsafe.Sizeof(p3)) // uint8 uint8 uint8 1 byte - это и есть byte

	const (
		p4 = iota
		q4 // какой тип переменной будет? какое будет значение?
		r4
	)
	fmt.Printf("%T %T %T", p4, q4, r4) // int int int
	fmt.Println(p4, q4, r4)            // 0 1 2
}
