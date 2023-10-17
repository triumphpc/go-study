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

func (p field) printVal() {
	fmt.Println(p.name)
}

func main() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	dataVal := []field{{"one"}, {"two"}, {"three"}}

	// Тут передача по ссылке значения, поэтому нет проблемы передачи значения в горутину
	for _, v := range data {
		go v.print()
	}

	time.Sleep(1 * time.Second)

	for _, v := range dataVal {
		// Тут передача не по ссыке в горутину, поэтому будет браться послденее проставленное значение
		go v.print()
	}

	time.Sleep(1 * time.Second)

	// А тут передача по значению, т.е. передается копия переменной в функцию, которая уже вызывает горутину
	for _, v := range dataVal {
		go func(v field) {
			v.print()
		}(v)
	}

	time.Sleep(1 * time.Second)
}
