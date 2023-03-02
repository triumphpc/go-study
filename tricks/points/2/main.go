package main

import "fmt"

func main() {
	x := []int{1, 3}
	y := []int{2, 4}

	ref := x //  Тут аллоцируется нова область памяти и базовый массив
	x = y

	fmt.Printf("%v %v %v", x, y, ref) // [2 4] [2 4] [1 3]

}
