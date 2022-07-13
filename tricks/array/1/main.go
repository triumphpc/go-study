package main

import "fmt"

func main() {
	//in := []string{"one", "two", "three"}

	out := make([]string, 0, 3)
	fmt.Printf("%v %v\n", len(out), cap(out))

	out = append(out, "one", "two", "three")
	fmt.Printf("%v %v\n", len(out), cap(out))

}
