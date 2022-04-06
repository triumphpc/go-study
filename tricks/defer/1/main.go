package main

import "fmt"

type toProcess struct {
	toReturn int
}

func (t *toProcess) process() {
	t.toReturn = 1
	fmt.Println(t.toReturn)
}

func (t *toProcess) testFunc() int {
	defer t.process()
	t.toReturn = 2
	return t.toReturn
}

func main() {
	var t = toProcess{}
	fmt.Println(t.testFunc()) // 1 \n 2
}
