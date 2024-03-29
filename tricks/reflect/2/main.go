package main

import (
	"fmt"
	"reflect"
)

type T struct {
	X int
	x int
}

func (T) M() {}
func (T) m() {}
func (T) K() {}

func main() {
	v := reflect.ValueOf(T{})
	fmt.Println(v) // {0 0}

	fmt.Println(v.NumField())               // 2
	fmt.Println(v.NumMethod())              // 2 - видит только экспортированные методы
	fmt.Println(reflect.TypeOf(T{}).Name()) // T

}
