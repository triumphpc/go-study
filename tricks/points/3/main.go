package main

import "fmt"

func main() {
	s := "qwe"
	ps := &s         // указатель на область
	b := []byte(*ps) // отдельная аллокация на значение области ps
	pb := &b         // указатель на область b

	s += "r"                          // s, ps = qwer
	*ps += "t"                        // s, ps = qwert
	*pb = append(*pb, []byte("y")[0]) // qwey

	fmt.Printf(*ps, string(*pb)) // qwert %!(EXTRA string=qwey)
}
