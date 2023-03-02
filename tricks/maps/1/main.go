// Golang Скомпилируется?
package main

import "fmt"

type A struct {
	name    string
	surname string
}

type B struct {
	slice []int
}

func main() {
	m1 := make(map[A]int)
	//m2 := make(map[B]int) // slice[]int not use equal == operator
	// в качестве ключа нельзя использовать несравниваемые типы

	fmt.Println(m1)

}
