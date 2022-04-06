package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data := []*field{{"one"}, {"two"}, {"three"}}

	// Тут передача по ссылке значения, поэтому нет проблемы передачи значения в горутину
	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
