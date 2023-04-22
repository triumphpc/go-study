package main

// Приведение слайса к массиву
// https://uproger.com/12-priyomov-go-kotorye-izmenyat-vashu-produktivnost/
import (
	"fmt"
	"testing"
)

func main() {
	//a := []int{0, 1, 2, 3, 4, 5}
	//var b [3]int = a[0:3] // // cannot use a[0:3] (value of type []int) as [3]int value in variable
	// declaration compiler(IncompatibleAssign)

	// go 1.20
	//a := []int{0, 1, 2, 3, 4, 5}
	//b := [3]int(a[0:3])
	//
	//fmt.Println(b) // [0 1 2]

	// go 1.17
	a := []int{0, 1, 2, 3, 4, 5}
	b := *(*[3]int)(a[0:3]) // Указатель на указатель

	fmt.Println(b) // [0 1 2]

}

// go 1.17
func TestM2e(t *testing.T /* Тут что-то*/) {
	a := []int{0, 1, 2, 3, 4, 5}
	b := *(*[3]int)(a[0:3])

	fmt.Println(b) // [0 1 2]
}
