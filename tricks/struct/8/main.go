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
			t.n = 3     // Тут копия значения в цикле
			ts[1].n = 9 // а тут установка конкретного значения
		case 1:
			fmt.Println(t.n, " ") // 9
		}
	}

	fmt.Println(ts) // [{0} {9}]

	ts = [2]T{}

	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3     // Тут копия значения в цикле
			ts[1].n = 9 //
		case 1:
			fmt.Println(t.n, " ") // 0, берет из значения из цикла
		}
	}

	fmt.Println(ts) // [{0} {9}]

	ts = [2]T{}

	for i := range ts[:] { // тут проходит по новому слайса
		switch t := &ts[i]; i { // Тут приравнивает к t для области t и проходит по i
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ")
		}
	}

	fmt.Println(ts) // [{3} {9}]

}
