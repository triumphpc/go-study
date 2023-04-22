package main

import "fmt"

type Test interface {
	Set(int)

	Get() int
}

type Imp struct {
	val int
}

type request struct {
}

func (i *Imp) Set(v int) {
	i.val = v

}

func (i *Imp) Get() int {
	return i.val
}

func main() {

	tt := ret()

	change(tt)

	fmt.Println(tt.Get())

	otherFun(tt)

	fmt.Println(tt.Get())

}

func ret() Test {
	return &Imp{}

}

func change(t Test) {
	t.Set(555)

}

func otherFun(t Test) {
	req := &request{}
	req.changeINOther(t)
}

func (t *request) changeINOther(tt Test) {
	tt.Set(66)
}
