package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f func(int) // Тут объявление ссылочного типа (функция)

	{ // Это просто скобочный блок видимости
		var f = func() {
			fmt.Printf("%T\n", f) // func(int) Тут вывод тип переменной просто
		}

		f()
	}

	type A = int // эта переменная переопределяется дальше

	{
		type A struct {
			*A
		}

		fmt.Println(reflect.TypeOf(A{}.A).Elem().Kind()) // struct Тут выводит ключевое слово для ссылочного типа A
	}
}
