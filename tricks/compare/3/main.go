// Для сравнения структур используем  DeepEqual
//
package main

import (
	"fmt"
	"reflect"
)

type data struct {
	num    int               // ok
	checks [10]func() bool   // несравниваемо
	doit   func() bool       // несравниваемо
	m      map[string]string // несравниваемо
	bytes  []byte            // несравниваемо
}

type S struct {
	a, b string
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2)) // prints: v1 == v2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) // prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2)) // prints: s1 == s2: true

	// Сравнение интерфейсов

	x := interface{}(&S{"a", "b"})
	y := interface{}(&S{"a", "b"})
	fmt.Println(reflect.DeepEqual(x, y)) // true
}
