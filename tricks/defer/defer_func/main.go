package main

import "fmt"

type toProcess struct {
	toReturn int
}

func (t *toProcess) process() {
	t.toReturn = 1
	fmt.Println(t.toReturn)
}

func (t *toProcess) testFunc() int {
	// 0
	defer t.process() // Выполнение при возврате, на момент вызова равен 2, на выходе 1, но не обрабатывается

	// 0
	t.toReturn = 2

	// 2
	return t.toReturn // Возвращаем 2
}

func (t *toProcess) testFunc2() int {
	// 0
	defer func() { // Выполняется после возврата
		// 2
		t.process()
		// 1
	}()

	// 0
	t.toReturn = 2

	// 2

	return t.toReturn
}

func main() {
	var t = toProcess{}
	fmt.Println(t.testFunc())  // 1 \n 2
	fmt.Println(t.testFunc2()) // 1 \n 2
}
