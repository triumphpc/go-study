package main

func main() {
	x := []int{1, 2, 3}

	y := [3]*int{}

	for i, v := range x {
		defer func() {
			println(v)
		}()

		y[i] = &v
	}

	println(*y[0], *y[1], *y[2]) // 3 3 3 3 3 3
}
