// Appending slices
package main

import "fmt"

func main() {
	var slice1 = make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	slice2 := slice1[2:4]
	inspectSlice(slice2) // ??

	slice2[0] = "CHANGED"
	inspectSlice(slice1) // ??
	inspectSlice(slice2) // ??

	slice2 = append(slice2, "NEW")
	inspectSlice(slice1) // ?
	inspectSlice(slice2) // ?

	slice3 := slice1[2:4:4]
	inspectSlice(slice3) // ??

	slice3 = append(slice3, "NEW 2")
	inspectSlice(slice1) // ?
	inspectSlice(slice3) // ?

	slice4 := make([]string, len(slice1))
	copy(slice4, slice1)

	inspectSlice(slice4) // ?

	slice4[0] = "UP"
	inspectSlice(slice4) // ?
	inspectSlice(slice1) // ?

}

func inspectSlice(s []string) {
	fmt.Printf("Length:[%v], ", len(s))
	fmt.Printf("Cap [%v]\n", cap(s))

	for i, v := range s {
		fmt.Printf("[%v] [%p] [%v]\n", i, &v, v)
	}
}
