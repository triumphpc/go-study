package main

import "fmt"

type Model interface {
	Test()
}

type Test struct {
}

func (t *Test) Test() {

}

func ret() Model {
	t := &Test{}

	return t
}

func main() {
	var emails []string

	fmt.Println(len(emails))
	//uniq := make(map[string]struct{}, 2)
	//uniq2 := make(map[string]bool, 2)
	//
	//uniq["test"] = struct{}{}
	//uniq["test2"] = struct{}{}
	//
	//uniq2["test"] = true
	//uniq2["test2"] = true
	//
	//fmt.Println(unsafe.Sizeof(uniq["test"]))
	//fmt.Println(unsafe.Sizeof(uniq2["test"]))
}

func Reverse(input []string) {

}
