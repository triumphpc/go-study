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

}
