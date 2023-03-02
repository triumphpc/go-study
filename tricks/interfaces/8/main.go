package main

import "fmt"

type S struct {
}

func g(x *interface{}) {

}

func main() {
	//g(S{}) //  cannot use S{} (value of type S) as type *interface{} in argument to g:
	//S does not implement *interface{} (type *interface{} is pointer to interface, not interface)
	s := S{}
	intS := interface{}(s)
	g(&intS)

	// Указатель на структуру и указатель на интерфейс — это не одно и то же.
	//Интерфейс может хранить либо структуру напрямую, либо указатель на структуру.
	//В последнем случае вы все равно просто используете интерфейс напрямую, а не указатель на интерфейс. Например:
	var f1 Foo
	var f2 *Foo = &Foo{}

	DoFoo(f1) //[main.Foo] {}
	DoFoo(f2) //[*main.Foo] &{}

	intF2 := interface{}(f2) //
	g(&intF2)

}

type Fooer interface {
	Dummy()
}

type Foo struct{}

func (f Foo) Dummy() {}

func DoFoo(f Fooer) {
	fmt.Printf("[%T] %+v\n", f, f)
}
