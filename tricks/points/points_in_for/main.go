package main

import (
	"fmt"
)

// Написать возведение в степень элементов слайса по указателю
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
		fmt.Printf("%d \n", *u) // 4 4 4 4
	}

	//a = []int{1, 2, 3}
	//for i := range a {
	//	result = append(result, &a[i])
	//}
	//
	double(result)

	for i := range result {
		fmt.Println(*result[i])
	}
}

// Как возвести в степень?
func double(arr []*int) {
	mm := make(map[*int]struct{}, len(arr))

	for i := range arr {
		if _, ok := mm[arr[i]]; ok {
			continue
		}

		mm[arr[i]] = struct{}{}

		*arr[i] = *arr[i] * 2
	}

}
