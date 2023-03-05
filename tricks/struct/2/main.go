package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	a = 3
	fmt.Println(unsafe.Sizeof(a)) // prints 8 bit

	var b uint16
	b = 3
	fmt.Println(unsafe.Sizeof(b)) // prints 2 byte 16 bit

	type S struct {
		a uint16
		b uint32
	}
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6

	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // prints 0 byte

	var x [1000000000]struct{}
	fmt.Println(unsafe.Sizeof(x)) // prints 0

	var x2 = make([]struct{}, 1000000000)
	fmt.Println(unsafe.Sizeof(x2)) // 24

	var x3 = make([]struct{}, 100)
	var y = x3[:50]
	fmt.Println(len(y), cap(y)) // prints 50 100

	var a3, b3 struct{}

	fmt.Printf("%p %p\n", a3, b3) // false  %!p(struct {}={}) %!p(struct {}={})

	a4 := make([]struct{}, 10)
	b4 := make([]struct{}, 20)
	fmt.Println(&a4 == &b4)       // false, a and b are different slices
	fmt.Println(&a4[0] == &b4[0]) // true, their backing arrays are the same

	a5 := struct{}{} // not the zero value, a real new struct{} instance
	b5 := struct{}{}
	fmt.Println(a5 == b5) // true

}
