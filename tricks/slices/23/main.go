package main

import "fmt"

func main() {
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := num[:5:8]
	fmt.Println(s1, len(s1), cap(s1)) // [0 1 2 3 4] 5 8

	//s2 := num[::4] // так нельзя объявлять
	//fmt.Println(s2)

	sl := make([]int, 0, 2)

	ll := []int{1, 2, 3}
	for _, v := range ll {
		sl = append(sl, v)
	}

	fmt.Println(sl, len(sl), cap(sl)) // [1 2 3] 3 4

}
