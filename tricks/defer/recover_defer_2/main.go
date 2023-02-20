package main

import "fmt"

// Вывод:
//2
//3 33
//1 22

func main() {
	defer func() {
		fmt.Println("1", recover()) // 1 и выводит панику 22
	}()

	defer func() {
		fmt.Println("2")                  // 2
		defer fmt.Println("3", recover()) // 3 и выводит панику 33
		panic(22)                         // Вызывает вторую понику
	}()

	defer recover() //Тут ничего не произойдет
	recover()       // Как и тут

	panic(33) // Тут начинает паниковать
}
