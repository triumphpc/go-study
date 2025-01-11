package main

import "log"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl1 := arr[1:4]
	sl2 := append(sl1, 10)
	sl2[0] = 99

	log.Println(arr) // 1
	log.Println(sl1) // 2
	log.Println(sl2) // 3
}
