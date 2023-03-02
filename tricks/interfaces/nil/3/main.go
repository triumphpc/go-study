package main

import (
	"fmt"
	"log"
	"reflect"
)

var x = *new(*int) // т.е. берем значение указателя на указатель типа int
var y *int = nil

func f() interface{} {
	return y
}

func main() {
	fmt.Printf("%T %v\n", x, x) // *int, nil - указатель
	log.Println(x == nil)       // true
	fmt.Printf("%T %v\n", y, y) // *int <nil>
	log.Println(y == nil)       // true

	fmt.Printf("%T %v\n", f(), f())          // *int <nil>
	fmt.Println(f() == nil)                  // false
	fmt.Printf("%s\n", reflect.ValueOf(f())) // %!s(*int=<nil>) - интерфейс, tab ссылка у которого указывает на y=*int
	fmt.Printf("%v\n", reflect.TypeOf(f()))  // *int указатель на тип int

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
