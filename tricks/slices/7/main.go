package main

import "fmt"

func main() {
	h, w := 2, 4

	raw := make([]int, h*w)
	for i := range raw {
		fmt.Println(i)
		raw[i] = i
	}
	// выводит: [0 1 2 3 4 5 6 7] <ptr_addr_x>

	table := make([][]int, h)
	for i := range table {
		// Тут получаем срез слайса [0:4][4:8]
		table[i] = raw[i*w : i*w+w]
	}

	fmt.Println(table, &table[1][0])
	// выводит: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>
}
