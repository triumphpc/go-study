package main

import "fmt"

func main() {
	s := "123"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	println(*ps)         // 12345
	println(string(*pb)) // 103

	*pb = append(*pb, []byte("y")[0])
	fmt.Println(string(*pb))

}
