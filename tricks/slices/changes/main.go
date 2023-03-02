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
	//testSlices9()
	testSlices10()

}

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2] // С 2 по 2 включительно

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b _] 1 2

	b[0] = "q"

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a q c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [q _] 1 2
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd') // Новый слайс, ссылается на тот же базовый массив,
	// которая ссылается на область базового массива, начиная с 1 элемента
	// Затем идет добавление элемента к слайсу и базовому массиву, получается:
	// a => a b d
	// b =>   b d

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b d] 2 2

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a z d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [z d] 2 2

	b = append(b, 'X')                                        // Аллоцируется новая область памяти и ссылка на новый базовый массив
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a z d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [z d X] 3 8

	b[0] = 'k'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a z d] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [k d X] 3 8
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

	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a d c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [a d]   2 2

	b[0] = 'z'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [z d c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) // [z d]   2 2
}

func testSlices7() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2:2], 'd') // Тут выделяется новый массив базовый и слайс. Слайс размер 1 и одним элементом
	// Затем сразу добавляется новый элемент и аллоцируется новый базовый массив, вместимостью 8
	c := a[1:2:2] // Тут ссылка на старый массив

	// Выделает новый массив большей длинны
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a b c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [b d] 2 8
	fmt.Printf("%p %s len=%v cap=%v\n", c, c, len(c), cap(c)) //   [b]   1 1

	b[0] = 'z'
	c[0] = 'x'
	fmt.Printf("%p %s len=%v cap=%v\n", a, a, len(a), cap(a)) // [a x c] 3 3
	fmt.Printf("%p %s len=%v cap=%v\n", b, b, len(b), cap(b)) //   [z d] 2 8
	fmt.Printf("%p %s len=%v cap=%v\n", c, c, len(c), cap(c)) //   [x]   1 1
}

func testSlices8() {
	var x = []int{4: 44, 55, 66, 1: 77, 88} // Тут просто выставляем значение для индекса
	//[]int{
	//	4: 44,
	//	55, // Продолжает последовательность после индекса
	//	66,
	//	1: 77,
	//	88,
	//}

	fmt.Println(x)        // [0 77 88 0 44 55 66]
	println(len(x), x[2]) // 7 88

}

func testSlices9() {
	arr := make([]int, 0, 10000)
	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr)) //  [] len=0 cap=10000

	arr = append(arr, 10)
	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr)) // [%!s(int=10)] len=1 cap=10000

	add(arr, 15) // 0xc0000b4000 [5 15] len=2 cap=10000

	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr)) // 0xc0000b4000 [5] len=1 cap=10000
	// Т.е. передается копия слайса. Копии слайса ничего не знают про cap, len оригинала:

	arr = make([]int, 0, 10000)

	arr2 := arr
	arr2 = append(arr2, 10)

	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr))     // 0xc00012a000 [] len=0 cap=10000
	fmt.Printf("%p %v len=%v cap=%v\n", arr2, arr2, len(arr2), cap(arr2)) // 0xc00012a000 [10] len=1 cap=10000

	// НО! если теперь мы в первом произведем добавление, то затрется эта 10ка
	arr = append(arr, 20)

	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr))     // 0xc0000ce000 [20] len=1 cap=10000
	fmt.Printf("%p %v len=%v cap=%v\n", arr2, arr2, len(arr2), cap(arr2)) // 0xc0000ce000 [20] len=1 cap=10000

}

func add(arr []int, v int) {
	arr[0] = 5
	arr = append(arr, v)
	fmt.Printf("%p %v len=%v cap=%v\n", arr, arr, len(arr), cap(arr)) // тут адрес один и тот же на базовый массив , но тут len 2

}

func testSlices10() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s1, cap(s1)) // [1 2 3 4 5 6 7 8] capacity = 8

	s2 := s1[:3]
	fmt.Println(s2, cap(s2)) // [1 2 3 _ _ _ _ _], capacity = 8

	s3 := s1[3:6]
	fmt.Println(s3, cap(s3)) // [      4 5 6 _ _], capacity = 5

	// append to s2
	fmt.Printf("\n\nappend 10 to s2\n")
	s2 = append(s2, 10)

	fmt.Println(s1, cap(s1)) // [1 2 3 10 5 6 7 8] 8
	fmt.Println(s2, cap(s2)) // [1 2 3 10 _ _ _ _] 8
	fmt.Println(s3, cap(s3)) // [      10 5 6 _ _] 5

	// // append to s1
	fmt.Printf("\n\nappend 20 to s1\n")
	s1 = append(s1, 20)

	fmt.Println(s1, cap(s1)) // [1 2 3 10 5 6 7 8 20] 16
	fmt.Println(s2, cap(s2)) // [1 2 3 10 _ _ _ _ _] 8
	fmt.Println(s3, cap(s3)) // [      10 5 6 _ _ _] 5

	// тут улетел массив для первого слайса в новую область памяти
	// и на нее теперь первый слайс ссылается,
	//при этом второй и третий ссылаюсят пока еще на старый массив

	s2[3] = 11
	fmt.Printf("\n\n")
	fmt.Println(s1, cap(s1)) // [1 2 3 10 5 6 7 8 20] 16
	fmt.Println(s2, cap(s2)) // [1 2 3 11 _ _ _ _ _] 8
	fmt.Println(s3, cap(s3)) // [      11 5 6 _ _ _] 5

	// // change value in s3
	fmt.Printf("\n\nchange s3[0]\n")
	s3[0] = 4

	fmt.Println(s1, cap(s1))
	fmt.Println(s2, cap(s2))
	fmt.Println(s3, cap(s3))
	// change s3[0]
	//[1 2 3 10 5 6 7 8 20] 16
	//[1 2 3 4] 8
	//[4 5 6] 5

}
