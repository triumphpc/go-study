package main

import "fmt"

// https://ueokande.github.io/go-slice-tricks/

func main() {
	// COPY
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make([]int, len(a))
	copy(b, a)                   // 1
	b = append([]int(nil), a...) // 2
	b = append(a[:0:0], a...)    // 3
	fmt.Println(a, b)

	// DELETE
	del := 5
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = append(a[:del], a[del+1:]...)
	fmt.Println(a, len(a), cap(a)) // 9 10

	// CUT
	i := 3
	j := 5
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy(a[i:], a[j:]) // [0 1 2 5 6 7 8 9 8 9] 10 10
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = 0 // or the zero value of T
	}
	//fmt.Println(a, len(a), cap(a)) // [0 1 2 5 6 7 8 9 0 0] 10 10
	a = a[:len(a)-j+i]
	fmt.Println(a, len(a), cap(a)) // [0 1 2 5 6 7 8 9] 8 10

	// FILTER
	n := 0
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	a = a[:n]

	fmt.Println(a) // [5]

}

func keep(x int) bool {
	return x == 5

}
