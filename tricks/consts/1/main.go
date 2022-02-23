package main

import "fmt"

const (
	_ = iota
	Avito
	There
	One
	Mmmm = iota
	Kkk
)

func main() {
	fmt.Println(Avito, There, One, Mmmm, Kkk)

}
