package main

import "fmt"

var x = *new(*int) // т.е. берем значение с ссылки на ссылку типа int
var y *int = nil

func f() interface{} {
	return y
}

func main() {
	fmt.Printf("%T\n", x)       // *int
	fmt.Printf("%T %v\n", y, y) // *int <nil>

	fmt.Printf("%T %v", f(), f()) // *int <nil>
	fmt.Println(f() == nil)       // false

	o := fmt.Print
	if f() == nil {
		if x == nil {
			o("A")
		} else {
			o("B")
		}
	} else {
		if x == nil {
			/// Тут будет вывод
		}

	}

}
