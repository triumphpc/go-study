package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v reflect.Value
	a := fmt.Sprint(v) == v.String()

	fmt.Println(a, fmt.Sprint(v)) // false <invalid reflect.Value>

	v = reflect.ValueOf(v)
	b := fmt.Sprint(v) == v.String()
	fmt.Println(b, fmt.Sprint(v)) // false <invalid Value>

	v = reflect.ValueOf(v)
	c := fmt.Sprint(v) == v.String()
	fmt.Println(c, fmt.Sprint(v)) // true <reflect.Value Value>

}
