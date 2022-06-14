// # github.com/triumphpc/go-study/tricks/maps/2
//./main.go:13:2: cannot assign to struct field m["x"].name in map

package main

type S struct {
	name string
}

func main() {
	//m := map[string]S{"x": S{"one"}}
	//fmt.Println("%v %p", m, m)
	//
	//m["x"].name = "test"
	//fmt.Println(m["x"].name)

}
