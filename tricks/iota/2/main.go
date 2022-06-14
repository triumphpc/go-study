package main

import "fmt"

const (
	Test  = "test"
	Test2 = 2

	state1 = iota + 1
	state2
	state3
)

const (
	fixeState = iota + 1
	fixStat2
)

func main() {

	fmt.Println(state2)   // 4
	fmt.Println(fixStat2) // 2

}
