package main

import "fmt"

type MyType int

type MyType2 struct {
	val []*MyType
}

type MyTypeStruct2 struct {
	name string
}

func main() {
	tt := MyType(1)

	var tt2 *MyType

	tt3 := MyType(55)
	tt2 = &tt3

	fmt.Println(tt)

	fmt.Println(tt2)

	test(tt)
	test(*tt2)

	fmt.Println("--------")

	testVal := MyTypeStruct2{
		name: "Тестовая",
	}

	var routes []*MyTypeStruct2

	routes = make([]*MyTypeStruct2, 0, 1)
	routes = append(routes, &testVal)

	fmt.Println(routes, cap(routes), len(routes))

	test2(routes)

}

func test(tt MyType) {
	fmt.Println(tt)

}

func test2(test []*MyTypeStruct2) {
	for id := range test {
		fmt.Println(test[id].name)
	}
}
