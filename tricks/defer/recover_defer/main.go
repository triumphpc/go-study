package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %#v", r)
		}
	}()

	panic(1) // recover in defer
	panic(2) // unreachable

}
func main() {
	f()
}
