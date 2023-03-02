// Что выведет?
package main

import "fmt"

type Test interface {
	string() string
}

type TestImpl struct {
	sentence string
}

func (stat *TestImpl) string() string {
	//return stat.sentence
	return ""

}

func main() {
	var t Test = &TestImpl{"Hello"}
	fmt.Println(t)

	var a int             // аллокация на типа int
	var b = &a            // b - это указатель на область памяти a
	fmt.Println(a, b, *b) //  выведет "0 0x10410020 0"

	var c int
	var d *int = &c   //  тоже самое, но с явным указанием типа указателя на число int
	fmt.Println(c, d) //  выведет "0 0x10410020"

}
