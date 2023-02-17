package main

import (
	"fmt"
	"time"
	"unsafe"
)

type MyType struct {
	//created time.Time
	updated *time.Time
}

func main() {
	var ttt *time.Time

	n := time.Now()
	ttt = &n

	fmt.Println(ttt)

	nn := new(MyType)
	nn.created = time.Now()

	fmt.Println(unsafe.Sizeof(nn))

}
