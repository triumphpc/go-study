package main

import "fmt"

type T1 struct{}

func (t1 T1) f() {
	fmt.Println("T1.f()")
}

func (t1 T1) g() {
	fmt.Println("T1.g()")
}

type T2 struct {
	T1
}

func (t2 T2) f() {
	fmt.Println("T2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	// Проверка интерфейса
	if t1, ok := i.(T1); ok {
		t1.f()
		t1.g()
	}

	if t2, ok := i.(T2); ok {
		t2.f()
		t2.g() //T1.g() так как не реализует интерфейс и тут метод ембеддинга
	}
}

func main() {
	printType(T1{})
	printType(T2{})
}
