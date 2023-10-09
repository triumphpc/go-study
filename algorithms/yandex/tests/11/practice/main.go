package main

import (
	"fmt"
)

func main() {
	digits := []int{10, 0, 7, 4, 3}

	fmt.Println(sortMerge(digits))

}

func sortMerge(digits []int) []int {
	if len(digits) <= 1 {
		return digits
	}

	middle := len(digits) / 2
	done := make(chan struct{})

	var left []int
	go func() {
		left = sortMerge(digits[:middle])
		done <- struct{}{}
	}()

	right := sortMerge(digits[middle:])

	<-done

	return merge(left, right)

}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged
}
