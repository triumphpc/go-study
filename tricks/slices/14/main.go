package main

import "fmt"

func rm(s []int, index int) []int {
	fmt.Println("тут 1")
	inspectSlice(s)
	arr := append(s[:index], s[index+1])

	fmt.Println("тут 2")
	inspectSlice(arr)

	return arr
}
func main() {
	var arr = []int{0, 1, 2, 3}
	inspectSlice(arr)

	for i := range arr {

		inspectSlice(arr)

		if arr[i] == 0 {
			inspectSlice(arr)
			//arr = rm(arr, i) //panic: runtime error: index out of range [1] with length 1

			//arr = append(arr[:i], arr[i+1]) // panic: runtime error: index out of range [1] with length 1
			//arr = []int{1} // panic: runtime error: index out of range [1] with length 1

			fmt.Println("тут 5")
			inspectSlice(arr)
		}
	}
}

func inspectSlice(s []int) {
	fmt.Printf("Length:[%v], ", len(s))
	fmt.Printf("Cap [%v]\n", cap(s))

	for i, v := range s {
		fmt.Printf("[%v] [%p] [%v]\n", i, &v, v)
	}
}
