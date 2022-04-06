package main

import "fmt"

func main() {
	//testSlices1()
	//testSlices2()
	//testSlices3()
	//testSlices4()
	//testSlices5()
	//testSlices6()
	//testSlices7()
	//testSlices8()

}

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2] // С 2 по 2 включительно

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b  ] 1 2

	b[0] = "q"

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a q c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [q  ] 1 2
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd') // Аллоцируется новая область памяти,
	// которая ссылается на область базового массива, начиная с 1 элемента
	// Затем идет добавление элемента к слайсу и базовому массиву, получается:
	// a => a b d
	// b =>   b d

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b d] 2 2

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a z d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [z d] 2 2
}

func testSlices3() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd', 'x')

	// Тут capacity больше базового размера массива, поэтому создается новый
	// базовый массив, ссылка на него уже. А остается ссылаться на старый

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [b d x _ _ _ _ _] 3 8

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [z d x _ _ _ _ _] 3 8
}

func testSlices4() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%s len=%v\n", b, len(b))                      // [a b c] 3
	a[0] = 'z'

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [z b c] 3 3
	fmt.Printf("%s len=%v\n", b, len(b))                      // [a b c] 3
}

func testSlices5() {
	//a := []byte{'a', 'b', 'c'}

	//b := append(a[1:2:1], 'd') // invalid index values, must be low <= high <= max
	//b[0] = 'z'
	//fmt.Printf("%s\n", a)
}

func testSlices6() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[0:1:2], 'd')
	// Ссылаются на одну и ту же область памяти
	// Создает новый слайс с ноым базовым массивом, который ссылается на первый элемент первоого массива
	// и имеет капасити 2, к нему добавлением еще один элемент

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [a d]   2 2

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [az b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [z d]    2 2
}

func testSlices7() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2:2], 'd')

	// Выделает новый массив большей длинны
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b d] 2 8

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [ab c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [z d] 2 8
}

func testSlices8() {
	var x = []int{4: 44, 55, 66, 1: 77, 88} // Тут просто выставляем знчение для индекса
	fmt.Println(x)                          // [0 77 88 0 44 55 66]
	println(len(x), x[2])                   // 7 88

}
