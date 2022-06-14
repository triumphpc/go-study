package main

import "fmt"

type S struct {
}

// Нельзя использовать указатель от другого типа
// Любой
func f(x *interface{}) {
	fmt.Printf("%v %T", x, x)
}
func g(x interface{}) {
	fmt.Printf("%v %T", x, x)

}

func main() {
	g(S{})
	g(&S{})
	//f(&S{}) // error

	x := interface{}(&S{})
	f(&x)

	var y interface{} = []int{1, 2, 3}
	f(&y)

}
