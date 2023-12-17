// Выравниванием памяти
// 32-битная система, вы можете получать 32-битные данные за один раз, то есть 4 байта, 64-битная - это 8 байтов, другими словами, в 64-битной системе метод хранения данных - 8-байтовое выравнивание

// Что указатель имеет постоянный размер и зависит от длины машинного слова, в данном случае это 8 байтов.
// Машинное слово - это максимальный размер чтения за один раз
// В языке Си используется термин «машинное слово», размер которого соответствует 4 или 8 байтам (в зависимости от разрядности процессора компьютера  —  32 или 64).
// А в Go используется «требуемое выравнивание». Его значение равно размеру памяти, требующемуся самому большому полю в структуре.
// То есть, если в структуре есть только поля int32, то «требуемое выравнивание» составит 4 байта. А если есть и int32, и int64 то 8 байтов.

// Пример:
// type Foo struct {
//	aaa [2]bool
//	bbb int64
//	ccc [2]bool
//}

// xx00 0000 [8 байт 64 бита] - машинное слово 8 байт, под первое поле выделяется памяти и заполняется 2 байта
// xxxx xxxx - под второе 8 байта
// xx00 0000 - 2 байта на третье поле

// https://prog-bytes.hashnode.dev/golang-structs-memory-allocation
//

package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T0 struct {
}

type T00 struct {
	aaa bool
	bbb int32
}

type T1 struct {
	a int64 // 8 байт (64 бит)
	b int32 // 4
	c int64
}

// Посмотрим как заполняется машинное слово

type T2 struct {
	aaa bool  // 1 по умолчанию. Но с заполнением - 4. (требуемое выравнивание)
	bbb int32 // 4 как максимальное в этой структуре.
	ссс bool  // 1 по умолчанию. Но с заполнением - 4.
}

type Foo struct {
	// xx00 xxxx
	// xx00
	// 12 byte
	aaa [2]bool // offset of bytes: 0
	bbb int32   // offset of bytes: 4
	ccc [2]bool // offset of bytes: 8
}

// Bar т.е. тут видно, что выделяется меньше памяти
type Bar struct {
	// xxxx xxxx
	// 8 byte
	aaa [2]bool // offset of bytes: 0
	ccc [2]bool // offset of bytes: 2
	bbb int32   // offset of bytes: 4
}

func main() {

	x := T0{}
	y := &T0{} // Тут аллокация памяти на 1 байт

	var takeInt32 int32                              // 32 bit
	fmt.Println(takeInt32, unsafe.Sizeof(takeInt32)) // 0 4 bytes

	fmt.Println(x, unsafe.Sizeof(x)) // {} 0
	fmt.Println(y, unsafe.Sizeof(y)) // &{} 8 byte

	x1 := T00{}  // 8
	y1 := &T00{} // 8

	fmt.Println(x, unsafe.Sizeof(x1)) // {} 8
	fmt.Println(y, unsafe.Sizeof(y1)) // &{} 8

	var t1 T1
	fmt.Println(t1, unsafe.Sizeof(t1)) // {0 0 0} 24

	atomic.AddInt64(&t1.c, 1)
	fmt.Println(t1, unsafe.Sizeof(t1)) // {0 0 1} 24

	ff := Foo{}
	bb := Bar{}

	fmt.Println(x, unsafe.Sizeof(ff)) // {} 12
	fmt.Println(y, unsafe.Sizeof(bb)) // &{} 8

	// Показываем смещение ячеек памяти для данных структур будет так Foo:
	// xx00
	// xxxx
	// xx00
	//
	//Выделяем 12 байта, смещение байт 0, 4, 8
	// Требуемое выравнивание 4 байта (т.е. мы можем выравнивать на 4 байта)
	//
	// Для Bar:
	// xxxx
	// xxxx
	// Выделяем 8 байт, смещение байт 0, 2, 4

	fmt.Printf("offsets of fields: aaa: %+v; bbb: %+v; ccc: %+v\n", unsafe.Offsetof(ff.aaa), unsafe.Offsetof(ff.bbb), unsafe.Offsetof(ff.ccc))
	// Требуемое выравнивание
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(ff))

	fmt.Printf("offsets of fields: aaa: %+v; ccc: %+v; bbb: %+v\n", unsafe.Offsetof(bb.aaa), unsafe.Offsetof(bb.ccc), unsafe.Offsetof(bb.bbb))
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(bb))

}
