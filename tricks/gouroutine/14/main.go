// Он запускается после сборки мусора, выражений go, операций блокирования каналов, блокирующих системных вызовов
//и операций блокирования. Также он может работать, когда вызвана невстроенная (non-inlined) функция.
// Чтобы узнать, встроена ли вызываемая вами в цикле функция, передайте -gcflags –m в go build или go run (например, go build -gcflags -m).

// Другое решение: явно вызвать диспетчер. Это можно сделать с помощью функции Gosched() из пакета runtime.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		fmt.Println("not done!") // не встроена
	}
	fmt.Println("done!")

	done = false

	go func() {
		done = true
	}()

	for !done {
		fmt.Println("not done!") // не встроена
		runtime.Gosched()
	}
	fmt.Println("done!")

}
