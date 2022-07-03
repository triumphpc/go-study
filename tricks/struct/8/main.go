package main

// В какой строке ошибка?

type T []int

func (T) X()  {}
func (*T) Z() {}

func main() {
	var t T

	t.X()
	t.Z()

	var p = &t
	p.X()
	p.Z()

	T{}.X()
	//T{}.Z()

}
