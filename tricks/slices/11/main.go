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
	a := []S{{1}, {5}, {6}}

	sort.Slice(a, func(i, j int) bool {
		return i < j

	})

	fmt.Println(a)

}
