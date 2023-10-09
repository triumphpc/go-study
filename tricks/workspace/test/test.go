package test

import "fmt"

type Test interface {
	First()
}

type FirstSt struct {
	Val string
}

func (f *FirstSt) First() {
	fmt.Printf("Struct 1 First %v\n", f.Val)
}
