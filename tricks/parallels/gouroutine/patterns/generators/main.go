package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Эта функция будет бесконечно повторять значения, которые вы ей передаете, пока вы не прикажете ей остановиться
	repeat := func(
		done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	// будет брать только первое число элементов из входящего потока valueStream, а затем завершится
	take := func(
		done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1, 2), 10) {
		fmt.Printf("%v ", num)
	}

	// Функция генерации случайного
	repeatFn := func(
		done <-chan interface{}, fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	done2 := make(chan interface{})
	defer close(done2)
	rand := func() interface{} {
		return rand.Int()
	}
	for num := range take(done2, repeatFn(done2, rand), 10) {
		fmt.Println(num)
	}

	// Преобразование входящего потока к строке
	toString := func(
		done <-chan interface{}, valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done3 := make(chan interface{})
	defer close(done3)
	var message string
	for token := range toString(done3, take(done3, repeat(done3, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
