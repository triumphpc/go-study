package main

import (
	"fmt"

	. "github.com/rogpeppe/fastuuid" //  Можно вот так импортировать, но не принято
	// Тогда работа будет относительно текущего пакета работа
)

func main() {
	g, _ := NewGenerator()
	fmt.Println(g)

}
