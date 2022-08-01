// go build -o bubble.out
// $ ./bubble.out -cpuprofile cpui.out -kind strinterface
// $ go tool pprof -list bubbleSortInterface ./bubble.out cpui.out

package main

import (
	"fmt"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

func main() {
	var ss []string
	for i := 0; i < 1000; i++ {
		ss = append(ss, "str"+strconv.Itoa(i))
	}

	bubbleSortGeneric(ss)
	fmt.Println("Done...")

}

func bubbleSortInterface(x sort.Interface) {
	n := x.Len()
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x.Less(i, i-1) {
				x.Swap(i, i-1)
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func bubbleSortGeneric[T constraints.Ordered](x []T) {
	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
