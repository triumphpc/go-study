package main

import "fmt"

type Test interface {
	string() string
}

type TestImpl struct {
	sent string
}

func (stat *TestImpl) string() string {
	return stat.sent
}

func main() {
	var t Test = &TestImpl{"test"}
	fmt.Printf("%p, %v", t, t)
	fmt.Println(t.string()) // test

}
