package main

import "fmt"

type S struct {
	name string
}

func main() {
	m := map[string]S{
		"x": S{"one"},
	}
	fmt.Printf("%v %p\n", m, m) // map[x:{one}] 0xc00007e180
	//
	//m["x"].name = "test" // Нельзя по ключу выставлять значения
	fmt.Println(m["x"].name) // one

	// Можно выставить через отдельную переменную
	nv := m["x"]
	nv.name = "new value"

	m["x"] = nv
	fmt.Println(m["x"].name) // new value

}
