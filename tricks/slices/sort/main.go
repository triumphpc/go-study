// Сортировка элементов слайса
package main

import (
	"errors"
	"fmt"
)

type S struct {
	v int
}

func main() {
	//	a := []S{5: {10}, 1: {5}, 0: {6}}
	//
	//	sort.Slice(a, func(i, j int) bool {
	//		return i < j
	//
	//	})
	//
	//	fmt.Println(a) // [{6} {5} {0} {0} {0} {10}]
	//
	//	ss := make([]uint, 0, 10)
	//
	//	ss = append(ss, 5, 4, 7, 8, 6, 1, 3, 9, 10)
	//
	//	slices.Sort(ss)
	//
	//	spentIdx := 1
	//	for _, s := range ss {
	//		log.Println(spentIdx, int(s))
	//
	//		if spentIdx == int(s) {
	//			spentIdx++
	//		} else {
	//			break
	//		}
	//	}
	//
	//	log.Println(spentIdx)
	fmt.Println(test())
}

func test() (err error) {

	defer func() {
		if err != nil {
			return
		}

	}()

	err = errors.New("NEW")

	return err

}
