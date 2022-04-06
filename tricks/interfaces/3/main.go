// Интерфейсный тип и поля значений заполняются на основе типа и значения переменной,
//использованной для создания соответствующей интерфейсной переменной. Если вы попытаетесь проверить,
//эквивалентна ли переменная nil, то это может привести к непредсказуемому поведению.
package main

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) // выводит: <nil> true
	fmt.Println(in, in == nil)     // выводит: <nil> true

	in = data
	fmt.Println(in, in == nil) // выводит: <nil> false
	//'data' является 'nil', но 'in' — не 'nil'

	doit := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		}

		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res) // выводит: good result: <nil>
		// 'res' не является 'nil', но его значение — 'nil'
	}

	doit2 := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil // возвращает явный 'nil'
		}

		return result
	}

	if res := doit2(-1); res != nil {
		fmt.Println("good result:", res)
	} else {
		fmt.Println("bad result (res is nil)") // здесь — как и ожидалось
	}
}
