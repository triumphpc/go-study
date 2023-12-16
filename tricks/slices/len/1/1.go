package main

import "fmt"

func main() {
	var hmm = [...]int{1, 2, 3, 12: 9, 8: 3}

	fmt.Println(len(hmm)) // 13
	// Выделяется память на количество выставленного индекса 12

	fmt.Println(hmm)

}
