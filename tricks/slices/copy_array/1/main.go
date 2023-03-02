// При присвоении массива создается точная копия данных

package main

import (
	"fmt"
)

func main() {
	// создает array с cap размером елментов
	a := [...]int{1, 2, 3, 4, 56}

	b := a

	fmt.Println(b == a) // true

	fmt.Printf("%v %p %T\n", a, &a, a)
	fmt.Printf("%v %p %T\n", b, &b, b) // it's different address

	a[0] = 42
	fmt.Printf("%v %p %T\n", b, b, b)

	fmt.Println(b == a) //false

}
