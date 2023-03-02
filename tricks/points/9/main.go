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
	tt2 = &tt3 // указатель

	fmt.Printf("%+v, %+v %T\n", tt, &tt, tt)    // 1, 0xc0000ac008 main.MyType
	fmt.Printf("%+v, %+v %T\n", tt2, *tt2, tt2) // 0xc0000ac010, 55 *main.MyType

	// передаем по значению
	test(tt) // 1, 0xc0000ac030 main.MyType
	// 33, 0xc000122030 main.MyType
	test(*tt2) // 55, 0xc0000ac038 main.MyType
	// 33, 0xc000122038 main.MyType

	fmt.Println("--------")

	testVal := MyTypeStruct2{
		name: "Тестовая",
	}

	var routes []*MyTypeStruct2

	routes = make([]*MyTypeStruct2, 0, 1)
	routes = append(routes, &testVal)

	fmt.Println(routes, cap(routes), len(routes)) // [0xc00010a210] 1 1
	test2(routes)                                 // Тестовая

}

func test(tt MyType) {
	fmt.Printf("%+v, %+v %T\n", tt, &tt, tt) // 0xc0000ac010, 55 *main.MyType
	tt = 33

	// Тут он выставляет
	fmt.Printf("%+v, %+v %T\n", tt, &tt, tt)

}

func test2(test []*MyTypeStruct2) {
	for id := range test {
		fmt.Println(test[id].name)
	}
}
