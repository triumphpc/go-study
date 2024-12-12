package main

import "log"

func main() {
	i := 0
	n := 10
	a := [11]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0} // 0 9 8 7 6 0 0 0 0 0 0

	for i < n {
		log.Println(a[i])
		log.Println(a)
		a[a[i]] = 0
		log.Println(a)

		i++
	}

	log.Println(a)

}
