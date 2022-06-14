package main

func main() {
	// Скомпилируется ли код?
	x := []int{123}
	x, x[0] = nil, 456
	x, x[0] = []int{555}, 456 // тут будет panic, потому что x стал nil
}
