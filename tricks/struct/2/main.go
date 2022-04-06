// Напрямую менять поле струкруты в слайсе нельзя
// Это не работает, потому что элементы таблицы не адресуемы.

// Первое обходное решение: использовать временную переменную.

package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v", m) //выводит: map[x:{two}]

	// Или менять по указателю
	m2 := map[string]*data{"x": {"one"}}
	m2["x"].name = "two" //ok
	fmt.Println(m["x"])  //выводит: &{two}

	//m3 := map[string]*data{"x": {"one"}}
	//m3["z"].name = "what?" //??? Ошибка компиляции

	//fmt.Println(m3)

}
