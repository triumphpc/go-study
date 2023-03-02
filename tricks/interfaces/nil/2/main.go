package main

import (
	"fmt"
	"log"
	"reflect"
)

type S struct{}

func initPointer() *S {
	var s *S

	return s
}

func initType() S {
	var s S

	return s
}

func initInterface() interface{} {
	var s S

	return s
}

type IF interface {
	F()
}

func (s S) F() {

}

func initInterfacePointer() interface{} {
	var s *S

	return s
}

func initInterfaceType() IF {
	var s S

	return s
}

func initInterfaceTypePointer() IF {
	var s *S

	return s
}

func main() {
	fmt.Println(initPointer() == nil) // true тут указатель на структуру - всегда nil

	fmt.Println(initType()) // {} аллокация пустого типа

	//fmt.Println(initType() == nil) // error, тип структуры нельзя сравниться с nil

	fmt.Println(initInterface() == nil)                  // false, так как поле tab интерфейса ссылается на S
	fmt.Printf("%s\n", reflect.ValueOf(initInterface())) // {}
	fmt.Printf("%v\n", reflect.TypeOf(initInterface()))  // main.S

	fmt.Println(initInterfacePointer() == nil)                  // false, аналогично
	fmt.Printf("%s\n", reflect.ValueOf(initInterfacePointer())) // %!s(*main.S=<nil>) - указатель на структуру main.S
	fmt.Printf("%v\n", reflect.TypeOf(initInterfacePointer()))  // *main.S - тип указателя на структуру

	fmt.Println(initInterfaceType() == nil)                  // false
	fmt.Printf("%s\n", reflect.ValueOf(initInterfaceType())) // {} Аналогично initInterface, т.е. возвращается интерфейсный тип,
	fmt.Printf("%v\n", reflect.TypeOf(initInterfaceType()))  //main.S - возвращаемый тип

	fmt.Println(initInterfaceTypePointer() == nil)                  // false
	fmt.Printf("%s\n", reflect.ValueOf(initInterfaceTypePointer())) // {} Аналогично initInterfacePointer
	fmt.Printf("%v\n", reflect.TypeOf(initInterfaceTypePointer()))

	var i interface{} // тут реально пустой интерфейс, нет базового тип на него ссылающегося

	if i == nil {
		fmt.Println("nil") // true
	}

	p := initPointer()
	log.Println(p == nil) // true
	i = p
	log.Println(i == nil)                  // false
	fmt.Printf("%s\n", reflect.ValueOf(i)) // %!s(*main.S=<nil>) - т.е. tab ссылается на базовый тип, который nil
	fmt.Printf("%v\n", reflect.TypeOf(i))  //*main.S - указатель на данный тип
}
