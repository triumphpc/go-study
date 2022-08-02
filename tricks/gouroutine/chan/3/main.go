package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var ch chan int
	//
	//<-ch // deadlock

	var i int
	fmt.Println(reflect.TypeOf(i).Elem().Size())

}
