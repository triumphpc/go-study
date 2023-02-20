package main

import "fmt"

func main() {
	val := 1

	switch val {
	case 1, 2: // Несколько значений
		fmt.Println(val)
	default:
		fmt.Println("XXX")

	}

	example1()

}

type int interface{}

func example1() {
	switch int("A") { // проверка
	case 1:
		if false {
			//fallthrough // тут ошибка будет, так как не в зоне видимости switch
		}
	case "A":
		fmt.Println("A") // A
		fallthrough
	case "B":
		fmt.Println("B") // B
	case "C":
		fmt.Println("C")
	default:
		fallthrough // проброс на следующий
	case false:

	}
}
