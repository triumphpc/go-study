package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := ['d']rune{} // Аллокация на 100 бит
	fmt.Println(a, 'd', len(a))

	fmt.Println(unsafe.Sizeof(a[0])) // 4 bytes

	var s string
	fmt.Println(unsafe.Sizeof(s)) // 16 bytes

	var i int8
	fmt.Println(unsafe.Sizeof(i)) // 1 byte

	var i32 int                     // its int64
	fmt.Println(unsafe.Sizeof(i32)) // 8 byte

	var r rune                    // alias int32 = 4 byte = 32 bit
	fmt.Println(unsafe.Sizeof(r)) // 4 byte

	var st struct{}
	fmt.Println(unsafe.Sizeof(st)) // 0 byte

	var b bool
	fmt.Println(unsafe.Sizeof(b)) // 0 byte

}
