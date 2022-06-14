package main

import "fmt"

// что выведет код

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ")
		}
	}

	fmt.Println(ts) // [{3} {9}]
}
