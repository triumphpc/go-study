// Какая строка приводит к ошибке компиляции
package main

import "fmt"

type T struct{ n int }

var x, y = T{}, T{9}

func main() {
	//[3]T{}[0] = y // error
	f := []T{2: x}
	f2 := [3]T{2: x}
	fmt.Printf("%v %T\n", f2, f2) // [{0} {0} {0}]

	fmt.Printf("%v %T\n", f, f) // [{0} {0} {0}]

	f[0] = y
	fmt.Printf("%v %T\n", f, f) //[{9} {0} {0}]

	f2[0] = y
	fmt.Printf("%v %T\n", f, f) //[{9} {0} {0}]

	[]T{2: x}[0] = y // Типа тоже самое, но без присвоения переменной и это норм

	m := map[int]T{}
	fmt.Printf("%v %T\n", m, m) // map[] map[int]main.T
	m[5] = y
	fmt.Printf("%v %T\n", m, m) // map[5:{9}] map[int]main.T
	map[int]T{}[5] = y          // Типа присвоение

	[]T{2: x}[0].n = 6 // В слайсе допустимо обращение к полю структуры, так как хранится указатель
	//map[int]T{}[5].n = 6 //  error. В mape так нельзя
}
