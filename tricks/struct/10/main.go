package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ") // 9
		}
	}

	fmt.Println(ts) // [{0} {9}]

	ts = [2]T{}

	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ") // 0
		}
	}

	fmt.Println(ts) // [{0} {9}]

}
