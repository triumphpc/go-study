package main

import "fmt"

// В какой строке будет паника?
func main() {
	var x *[9]int
	fmt.Println(x, len(x), cap(x)) // nil, 9, 9
	x = &[...]int{8: 0}
	fmt.Println(x, len(x), cap(x)) //&[0 0 0 0 0 0 0 0 0] 9 9
	y := x[1:3]                    // тут получается capacity уменьшается на срез
	fmt.Println(y, len(y), cap(y)) // [0 0] 2 8
	_ = y[5:7]                     // Это норм
	fmt.Println(y, len(y), cap(y)) // [0 0] 2 8
	//_ = y[5:] // panic: runtime error: slice bounds out of range [5:2]
	// Когда  указывается правая часть среза, по-умолчанию она равна len, и если это больше левой границы - то ошибка

	x = nil
	for _, _ = range x {
	}

	_ = len(y) / len(x)

}
