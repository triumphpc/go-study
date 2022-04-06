// qwert

package main

import "fmt"

func main() {
	s := "qwe"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "r"
	*ps += "t"
	*pb = append(*pb, []byte("y")[0])

	fmt.Printf(*ps)
}
