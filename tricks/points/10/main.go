package main

import (
	"fmt"
	"time"
	"unsafe"
)

type MyType struct {
	created time.Time
	updated *time.Time
}

func main() {
	var ttt *time.Time

	n := time.Now()
	ttt = &n

	fmt.Println(ttt) // 2023-02-25 12:41:08.598364 +0300 MSK m=+0.000068617

	nn := new(MyType)
	nn.created = time.Now()

	fmt.Println(unsafe.Sizeof(nn)) // 8 байт

	hh := struct{}{}
	fmt.Println(unsafe.Sizeof(hh)) // 0 байт

	var ii int
	fmt.Println(unsafe.Sizeof(ii)) // 8 байт

}
