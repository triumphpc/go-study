package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f func(int)
	{
		var f = func() {
			fmt.Printf("%T", f) // func(int)
		}
		f()
	}

	type A = int
	{
		type A struct {
			*A // А тут ссылается сам на себя, так как область видимости не func
		}

		fmt.Println(
			reflect.TypeOf(A{}.A).Elem().Kind()) // struct
	}

}
