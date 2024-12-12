package main

import (
	"log"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}

	log.Println(s)
	result := slices.Concat(s, s2)
	log.Println(result)

}
