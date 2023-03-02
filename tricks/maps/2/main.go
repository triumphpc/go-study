package main

import "fmt"

func main() {
	inc := map[int]bool{1: false, 3: true}

	val, ok := inc[555]
	fmt.Println(val, ok) // false (дефолтное значение), false

	inc2 := map[int]int{1: 3, 3: 5}
	val2, ok2 := inc2[555]
	fmt.Println(val2, ok2) // 0 (дефолтное значение), false

}
