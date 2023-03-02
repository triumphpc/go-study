package main

import (
	"fmt"
)

func main() {
	testData2()
}

func testData2() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v // Тут ссылка на операнд цикла, поэтому все приведет к 4
	}
	for _, u := range result {
		fmt.Printf("%d ", *u) // 4 4 4 4
	}
}
