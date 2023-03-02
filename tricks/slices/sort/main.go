// Сортировка элементов слайса
package main

import (
	"fmt"
	"sort"
)

type S struct {
	v int
}

func main() {
	a := []S{5: {10}, 1: {5}, 0: {6}}

	sort.Slice(a, func(i, j int) bool {
		return i < j

	})

	fmt.Println(a) // [{6} {5} {0} {0} {0} {10}]

}
