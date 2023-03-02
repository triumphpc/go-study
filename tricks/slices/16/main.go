package main

import (
	"fmt"
	"log"
)

// В какой строке будет паника?
func main() {
	var x *[9]int
	fmt.Println(x, len(x), cap(x)) // <nil> 9 9

	x = &[...]int{8: 1}
	fmt.Println(x, len(x), cap(x)) // &[0 0 0 0 0 0 0 0 1] 9 9

	x[0] = 0
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4
	x[5] = 5
	x[6] = 6
	x[7] = 7
	x[8] = 8

	fmt.Println(x, len(x), cap(x)) // &[0 1 2 3 4 5 6 7 8] 9 9

	y := x[1:3]                    // тут получается capacity уменьшается на срез
	fmt.Println(y, len(y), cap(y)) // [1 2] 2 8 [1 2 x x x x x x]

	z := y[5:7]                    // Это норм
	fmt.Println(z, len(z), cap(z)) // [6 7] 2 3 [6 7 x] тут ссылается на базовый массив через слайс, индекс
	// используется базового массива &[x x x x x x [6 7 x]] 2 3 - капасити от базового

	fmt.Println(y, len(y), cap(y)) // [1 2] 2 8
	d := y[5:6]                    // от базового
	fmt.Println(d, len(d), cap(d)) // [6] 1 3 // [x x x x x [6 x x]]

	//_ = y[5:]                      // panic: runtime error: slice bounds out of range [5:2]
	// Когда  указывается правая часть среза, по-умолчанию она равна len, и если это больше левой части(5:) границы - то ошибка

	h := x[1:3]
	fmt.Println(h, len(h), cap(h)) // [1 2] 2 8

	//_ = h[5:]

	log.Println(len(x)) // 9
	x = nil             // указатель на массив не меняет ничего

	log.Println(len(x), cap(x)) // 9 9
	log.Println(len(y), cap(y)) // 2 8

	for _, _ = range x {
	}

	_ = len(y) / len(x)

}
