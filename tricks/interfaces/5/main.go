package main

func f(...interface{}) {

}

func main() {
	// Какая строка приводит к ошибке? Ответ - никакая
	f()
	f(nil)
	f(1, 2, 3)
	f([]int{1, 2, 3})
}
