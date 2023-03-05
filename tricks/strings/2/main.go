package main

import "fmt"

func main() {
	st := "string"
	fmt.Println(st[0])        // 115
	fmt.Printf("%T\n", st[0]) // выводит uint8
	// string unmutable type
	//st[0] = 4

	s := []byte(st)
	s[0] = 'X'
	fmt.Println(string(s)) //Xtring

}
