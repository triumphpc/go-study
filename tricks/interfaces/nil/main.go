// https://habr.com/ru/post/597461/
package main

import (
	"fmt"
	"reflect"
)

// создадим некоторый тип Человек,
// который умеет спать, есть и работать
type Man struct{}

func (m Man) Sleep() {}
func (m Man) Eat()   {}
func (m Man) Work()  {}

// и тип Пёс, который у нас умеет только лаять
type Dog struct{}

func (d Dog) Bark() {}

// теперь мы хотим понять, кто может стать программистом,
// для этого мы определяем интерфейс Programmer,
// и определяем в нём метод Work
// этим мы ограничиваем количество объектов,
// которые смогут быть программистами
// так случилось, что в нашем коде, чтобы стать программистом,
// достаточно уметь работать :)
type Programmer interface {
	Work()
}

func main() {
	var man *Man
	fmt.Printf("%#v\n", man) // (*main.Man)(nil)
	var worker Programmer
	fmt.Printf("%#v\n", worker) // <nil>

	worker = man
	fmt.Printf("%#v\n", worker) // (*main.Man)(nil)
	// наш воркер изменил тип, он больше не совсем nil,
	// а nil от типа *main.Man
	// что же это для нас значит, давайте смотреть

	fmt.Printf("%t\n", man == nil) // true
	// отлично, наш man равен nil,
	// мы же туда ничего не положили и ссылка равна nil
	fmt.Printf("%t\n", worker == man) // true
	// ну тоже хорошо, мы положили в переменную
	// worker переменную man,
	// и как мы уже убеждались ранее, они равны

	// ну и понятно, по законам математики,
	// если x = 0, а y = x, значит y = 0
	fmt.Printf("%t\n", worker == nil) // false
	// а вот и нет, это вам Go, а не математика

	fmt.Println("------")

	var worker2 Programmer
	fmt.Printf("%s\n", reflect.ValueOf(worker2)) // <invalid reflect.Value>
	// рефлект не может получить значение из переменной worker,
	// потому что его просто нет,
	// то есть поле `data` у нашего интерфейса не задано
	fmt.Printf("%v\n", reflect.TypeOf(worker2)) // <nil>
	// как не может получить и тип,
	// так как он тоже не задан в поле `tab`
	fmt.Printf("%t\n", worker2 == nil) // true
	// worker равен nil

	var man2 *Man
	fmt.Printf("%t\n", man2 == nil) // true
	// переменная man равняется nil
	worker2 = man2

	fmt.Printf("%s\n", reflect.ValueOf(worker2).Elem())
	// <invalid reflect.Value>
	// после присвоения переменной worker переменной man
	// мы точно также не можем получить значение через reflect
	// так как man у нас не содержит каких-либо данных, он равен `nil`
	fmt.Printf("%v\n", reflect.TypeOf(worker2)) // *main.Man
	// а вот с типом интереснее мы из переменной worker
	// получили тип переменной man, то есть *main.Man
	// соответственно поле `tab` у нашего интерфейса уже явно не пустое
	// и поэтому переменная worker больше не будет равна nil
	fmt.Printf("%t\n", worker2 == nil) // false

	fmt.Println("----")
	// для начала нам нужно инициализировать переменную,
	// которая будет принимать результат из функции raiseError
	var err error

	// пробуем вернуть из нашей функции ошибку
	err = raiseError(true)
	fmt.Println(err != nil)                         // true
	fmt.Printf("%v\n", reflect.TypeOf(err))         // *main.CustomError
	fmt.Printf("%#v\n", err)                        // &main.CustomError{}
	fmt.Printf("%v\n", reflect.ValueOf(err).Elem()) // CustomError
	// отлично, err у нас действительно не равен nil

	// а теперь вызовем функцию таким образом, чтобы она вернула nil
	err = raiseError(false)
	fmt.Println(err != nil)                         // true
	fmt.Printf("%v\n", reflect.TypeOf(err))         // *main.CustomError
	fmt.Printf("%#v\n", err)                        // (*main.CustomError)(nil)
	fmt.Printf("%v\n", reflect.ValueOf(err).Elem()) // <invalid reflect.Value>
	// после преобразования в интерфейс, наш nil уже не очень то и nil

}

// создаём свой собственный тип ошибки,
// обычно это бывает нужно чтобы передавать дополнительные данные
// вместе с ошибкой, так например сделано в библиотеке twirp
type CustomError struct{}

// этот метод нам нужен, чтобы реализовать интерфейс error
func (CustomError) Error() string { return "CustomError" }

// а в этой функции мы будем возвращать нашу ошибку, либо nil
func raiseError(raise bool) *CustomError {
	if raise == true {
		return &CustomError{}
	}
	return nil
}
