package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v reflect.Value

	fmt.Printf("%#T %#v\n", v, v)                   // reflect.Value <invalid reflect.Value>s
	fmt.Printf("%#T %#v\n", v.String(), v.String()) // string "<invalid Value>"

	a := fmt.Sprint(v) == v.String()
	fmt.Printf("%#T, %#v\n", a, a) // bool false

	v = reflect.ValueOf(v)

	b := fmt.Sprint(v) == v.String()
	fmt.Printf("%#T %#v\n", v, v)                   // Pointer)(nil), flag:0x0}
	fmt.Printf("%#T %#v\n", v.String(), v.String()) // string "<reflect.Value Value>"

	v = reflect.ValueOf(v)
	fmt.Printf("%#T %#v\n", v, v) // reflect.Value reflect.Value{typ:(*reflect.rtype)(0x10a29e0), ptr:(unsafe.Pointer)(0xc00000c078), flag:0x99}

	fmt.Printf("%#T %#v\n", v.String(), v.String()) // string "<reflect.Value Value>"

	c := fmt.Sprint(v) == v.String()

	fmt.Println(a, b, c) // false false true

}
