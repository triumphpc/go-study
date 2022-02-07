package main

import (
	"errors"
	"fmt"
)

func main() {
	list := []int{1, 3, 5, 7, 9}

	res, err := lineSearch(list, 3)

	fmt.Println(res, err)

}

// Implementation line algorithm
func lineSearch(list []int, item int) (int, error) {
	low := 0
	high := len(list) - 1

	for low < high {
		mid := low + high
		guess := list[mid]

		if guess == item {
			return mid, nil
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("unknown")
}
