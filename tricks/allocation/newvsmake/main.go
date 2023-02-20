// Выражение new(T) создает неименованную переменную типа Т, инициализирует ее нулевым значением типа Т и возвращает ее адрес, который представляет собой значение типа *Т

// make используется для инициализации и аллокации slice s, map s и channel s.
// make не только открывает память, но также инициализирует ее нулевое значение для типа памяти. Он также работает с new

package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	name string
}

func main() {
	f := new(MyType)
	i := new(int)

	fmt.Println(f)  // &{}
	fmt.Println(*f) // {}

	fmt.Println(*i) // 0

	f.name = "new value"
	fmt.Println(f) // &{new value}
	fmt.Println(*f)

	*i = 3
	fmt.Println(*i) // 3

	v1 := newVal()
	fmt.Println(*v1) // 0

	v2 := newVal2()
	fmt.Println(*v2) // 0

	st := &MyType{name: "test"}
	fmt.Println(st) // &{test}

	st2 := new(MyType)
	*st2 = MyType{name: "test2"}
	fmt.Println(st2) // &{test2}

	// make
	av := make([]int, 5)
	fmt.Printf("av: %p %#v , %v\n", &av, av, unsafe.Sizeof(av)) //av: 0xc0000a0018 []int{0, 0, 0, 0, 0} , 24
	av2 := make([]int, 0, 5)
	fmt.Printf("av2: %p %#v %v \n", &av2, av2, unsafe.Sizeof(av2)) //av2: 0xc0000a0048 []int{} 24
	av[0] = 1
	fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000046400 []int{1, 0, 0, 0, 0}
	mv := make(map[string]string)
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[string]string{}
	mv["m"] = "m"
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[string]string{"m":"m"}
	chv := make(chan string)
	fmt.Printf("chv: %p %#v \n", &chv, chv) //chv: 0xc000074028 (chan string)(0xc00003e060)
	go func(message string) {
		chv <- message // Save message
	}("Ping!")
	fmt.Println(<-chv) // Cancel Interest/"Ping!"
	close(chv)

}

func newVal() *int {
	return new(int)
}

func newVal2() *int {
	var i int
	return &i
}
