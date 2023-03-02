// Пока значение адресуемо (addressable), к нему можно применять метод, принимающий значение по указателю.
// Иными словами, в некоторых случаях вам не нужно иметь версию метода, принимающего параметр по значению.
//
// Но не каждая переменная адресуема. Элементы хеш-таблицы (map) неадресуемы. Переменные,
// на которые ссылаются через интерфейсы, тоже неадресуемы.
package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{"one"}
	d1.print() //ok

	//d2 := data{"two"}
	//var in printer = d2 // ошибка, через интерфейс неадресуемый
	//in.print()

	//m := map[string]data {"x":data{"three"}}
	//m["x"].print() //ошибка, в мапе неадресуемые поля
}
