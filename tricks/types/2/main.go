package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var x int  // 8 байт 64 бит
	var y int8 // 1 байт = 8 бит

	fmt.Println(x, unsafe.Sizeof(x))
	fmt.Println(y, unsafe.Sizeof(y))

}
