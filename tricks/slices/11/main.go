package main

import "fmt"

func main() {
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := num[:5:8] // capacity
	//s2 := num[::3] // error
	s3 := num[:3:2]

	fmt.Println(len(s1), cap(s1)) // 0 1 2 3 4 x x x [5, 8]
	//fmt.Println(len(s2), cap(s2)) // 0 1 2 3 4 x x x [5, 8]
	fmt.Println(len(s3), cap(s3))

}
