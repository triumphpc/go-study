package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f func(int) // Тут объявление ссылочного типа

	{ // Это просто скобочный блок
		var f = func() {
			fmt.Printf("%T", f) // Тут вывод тип переменной просто

		}

		f()
	}

	type A = int // эта переменная переопределяется дальше

	{
		type A struct {
			*A
		}

		fmt.Println(reflect.TypeOf(A{}.A).Elem().Kind()) // тут выводит ключевое слово для ссылочного типа A
	}
}
