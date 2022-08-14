package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []string
	b := []string{} // тут происходит аллокация уже

	fmt.Printf("%T %T", a, b)

	var (
		aa []byte
		bb []byte
	)

	aa, _ = json.Marshal(a)
	bb, _ = json.Marshal(b)

	fmt.Println(string(aa), string(bb))

}
