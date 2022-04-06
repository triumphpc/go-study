package main

import "fmt"

func main() {
	var data interface{} = "great"

	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", data)
	} else {
		fmt.Println("[not an int] value =>", data)
		//выводит: [not an int] value => 0 (not "great")
	}

	var data2 interface{} = "great"

	if res, ok := data2.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data2)
		// выводит: [not an int] value => great (as expected)
	}
}
