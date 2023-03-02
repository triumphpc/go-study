package main

import "fmt"

func main() {
	s1 := make([]int, 1)
	z1 := make([]int, 1, 2)
	fmt.Printf("%v %T %p\n", s1, s1, s1)

	s2 := s1 // ссылочный тип, это указывает на один базовый массив
	s2[0] = 1

	z2 := z1
	z2[0] = 1

	fmt.Printf("%v %T %p\n", s1, s1, s1) // [1] []int 0xc00001c098
	fmt.Printf("%v %T %p\n", s2, s2, s2) // [1] []int 0xc00001c098

	fmt.Printf("%v %T %p\n", z1, z1, z1) // [1] []int 0xc00001c0a0
	fmt.Printf("%v %T %p\n", z2, z2, z2) // [1] []int 0xc00001c0a0

	// create new address and new pointer
	s1 = append(s1, 44)
	z1 = append(z1, 55)

	fmt.Printf("%v %T %p\n", s1, s1, s1) // [1 44] []int 0xc00001c0d0
	fmt.Printf("%v %T %p\n", s2, s2, s2) // it's old value

	fmt.Printf("%v %T %p\n", z1, z1, z1) // [1 55] []int 0xc00001c0a0
	fmt.Printf("%v %T %p\n", z2, z2, z2) // [1] []int 0xc00001c0a0

	s2[0] = 5
	z2[0] = 5

	fmt.Printf("%v %T %p\n", s1, s1, s1) // [1 44] []int 0xc00001c0d0
	fmt.Printf("%v %T %p\n", s2, s2, s2) // [5] []int 0xc00001c098

	fmt.Printf("%v %T %p\n", z1, z1, z1) // [5 55] []int 0xc00001c0a0
	fmt.Printf("%v %T %p\n", z2, z2, z2) // [5] []int 0xc00001c0a0

	// Т.е. при создании нового слайса он ссылается на базовый массив и вс лучае добавление элементов
	// в рамках капасити, на ссылаются на базовый.
	// Если же добавление идет с увеличением вместимости, но создается новый базовый массив

	z2 = append(z2, 66)
	z2[0] = 6
	fmt.Printf("%v %T %p\n", z1, z1, z1) // [6 66] []int 0xc00001c0a0
	fmt.Printf("%v %T %p\n", z2, z2, z2) // [6 66] []int 0xc00001c0a0

	z2 = append(z2, 77)
	z2[0] = 7

	// Если происходит аллокация, создается новый базовый массив для расширяемого слайса
	fmt.Printf("%v %T %p\n", z1, z1, z1) // [6 66] []int 0xc00001c0a0
	fmt.Printf("%v %T %p\n", z2, z2, z2) // [7 66 77] []int 0xc00001e080

}
