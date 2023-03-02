package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []string  // nil
	b := []string{} // тут происходит аллокация уже

	fmt.Printf("%T %T\n", a, b)     // []string []string
	fmt.Println(a == nil, b == nil) // true, false

	var (
		aa []byte
		bb []byte
	)

	aa, _ = json.Marshal(a)
	bb, _ = json.Marshal(b)

	fmt.Println(string(aa), string(bb)) // null []

}
