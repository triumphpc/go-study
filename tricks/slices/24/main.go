package main

import "fmt"

type My struct {
	id int
}

func main() {
	arr := []*My{{1}, {2}, {3}, {4}, {5}}

	for id := range arr {
		if arr[id].id == 5 {
			arr[id] = arr[len(arr)-1]
			arr[len(arr)-1] = nil
			arr = arr[:len(arr)-1]

			break
		}
	}

	fmt.Println(arr, cap(arr), len(arr))

}
