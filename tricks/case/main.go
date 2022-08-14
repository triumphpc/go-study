package main

import "fmt"

func main() {
	val := 1

	switch val {
	case 1,
		2:
		fmt.Println(val)
		fmt.Println(val)
	default:
		fmt.Println("XXX")

	}

}

type int interface{}

func example1() {
	switch int("A") {
	case 1:
		if false {
			fallthrough // тут ошибка будет, так как не в зоне видимости switch
		}
	default:
		fallthrough
	case false:

	}
}
