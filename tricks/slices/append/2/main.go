package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	fmt.Printf("%v %v\n", len(ss), cap(ss)) // 3 4

	ss = append(ss, 4) // тут новый базовый массив

	fmt.Printf("%v %v\n", len(ss), cap(ss)) // 3 4

	for _, v := range ss {
		v += 10
	}

	fmt.Printf("%v %v\n", len(ss), cap(ss)) // 3 4

	for i := range ss {
		ss[i] += 10
	}

	fmt.Printf("%v\n", s) // 1 2 3

}
