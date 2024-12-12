package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// https://habr.com/ru/companies/vk/articles/776766/

func main() {

	var x uint16
	x = 32768

	PrintAsBinary(x)

	var b1, b2 [257]byte

	e1 := unsafe.Pointer(&b1)
	e2 := unsafe.Pointer(&b2)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println("Size of b1:", unsafe.Sizeof(b1))
	fmt.Println("Size of b2:", unsafe.Sizeof(b2))
	fmt.Println("Distance b1 to b2:", uintptr(e2)-uintptr(e1))

	// Указатели
	y := 123

	fmt.Println("pointer in main", &y)
	p := pointTest(&y)
	fmt.Println(y)
	fmt.Println("Distance main to func:", uintptr(unsafe.Pointer(&y))-p)

	p = pointTest2(&y)
	fmt.Println(y) // 321

	// Структуры
	// Преобразованрие структуры в массив байтов
	type struct4 struct {
		b1 byte
		b2 byte
		i  int16
	}

	z := struct4{
		b1: 2,
		b2: 4,
		i:  261,
	}

	fmt.Println(*(*[4]byte)(unsafe.Pointer(&z))) // [2 4 5 1]

	// Выравнивание структур
	type alignmentStruct12 struct {
		b1 byte
		i  uint32
		b2 byte
	}

	x1 := alignmentStruct12{
		b1: 255,
		i:  4294967295,
		b2: 255,
	}

	fmt.Println(unsafe.Sizeof(x1))                 // 12
	fmt.Println(*(*[12]byte)(unsafe.Pointer(&x1))) // [255 0 0 0 255 255 255 255 255 1 0 0]

	type alignmentStruct8 struct {
		b1 byte
		b2 byte
		i  uint32
	}
	x2 := alignmentStruct8{
		b1: 255,
		b2: 255,
		i:  4294967295,
	}

	fmt.Println(unsafe.Sizeof(x2))                // 8
	fmt.Println(*(*[8]byte)(unsafe.Pointer(&x2))) // [255 255 25 4 255 255 255 255]

	// Интерфейсы
	// : указатель на тип и указатель на само значение.
	type iface struct {
		t     unsafe.Pointer
		value unsafe.Pointer
	}

	var x3 any

	x3 = 10

	i := *(*iface)(unsafe.Pointer(&x3))
	v := *(*int)(i.value) // А затем нехитрыми манипуляциями с конвертацией пойнтера в инт-пойнтер и разыменованием получаем значение.

	fmt.Println(i)
	fmt.Println(v) // 10

	// Слайс и его устройство
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	x4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := unsafe.Pointer(&x4)
	fmt.Println(*(*slice)(s1)) // Видим адрес массива в памяти, длину и вместимость, равную девяти. Легко выводим значения массива:
	fmt.Println(*(*[9]int)((*(*slice)(s1)).array))

	x5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x6 := x5[3:6]

	s2 := unsafe.Pointer(&x5)
	fmt.Println("x1 pointer", s2)
	fmt.Println("x1 slice", *(*slice)(s2))
	fmt.Println("x1 array", *(*[9]int)((*(*slice)(s2)).array))

	s3 := unsafe.Pointer(&x6)
	fmt.Println("x2 pointer", s3)
	fmt.Println("x2 slice", *(*slice)(s3))
	fmt.Println("x2 array", *(*[6]int)((*(*slice)(s3)).array))

}

// Функция небольшая и универсальная, позволяет напечатать любую переменную как набор байт разделенных пробелом. Что довольно удобно для чтения. Важной частью тут является преобразование интерфейса к структуре iface, детальнее про это можно прочитать ниже, в разделе «Интерфейсы». Таким образом мы получаем указатель на участок памяти где находятся данные.
// Затем, с помощью reflect извлекаем размер полученной переменной и в цикле итерируемся по нему. В самом же теле преобразуем указатель на участок в простой байт и выводим его в двоичном формате. В конце просто смещаем указатель на размер байта.
func PrintAsBinary(a any) {
	type iface struct {
		t, v unsafe.Pointer
	}
	p := uintptr((*(*iface)(unsafe.Pointer(&a))).v)

	t := reflect.TypeOf(a)

	for i := 0; i < int(t.Size()); i++ {
		n := *(*byte)(unsafe.Pointer(p))
		fmt.Printf("%08b ", n)
		p += unsafe.Sizeof(n)
	}

	fmt.Print("\n")
}

// go:noinline
func pointTest(t *int) uintptr {
	fmt.Println("pointer of t", &t)
	fmt.Println("pointer in func before", t)
	z := 321
	t = &z
	fmt.Println("pointer in func after", t)
	return uintptr(unsafe.Pointer(&t))
}

func pointTest2(t *int) uintptr { // Тут уже раизменовываем происходит
	fmt.Println("pointer of t", &t)
	fmt.Println("pointer in func before", t)
	z := 321
	*t = z // Разименовывание это обращение к значению по указателю переменной
	fmt.Println("pointer in func after", t)
	return uintptr(unsafe.Pointer(&t))
}
