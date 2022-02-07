// Example how use CG with slices and pointers

package main

import (
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var n = 40000000
	var structure []data

	for i := 0; i < n; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()

	_ = structure[0]

}
