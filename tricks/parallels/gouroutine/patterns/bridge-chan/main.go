package main

import "fmt"

// Функция, которая может преобразовать канал каналов в простой канал - метод, называемый соединением каналов, -
// это значительно упростит потребителю сосредоточение внимания на текущей проблеме. Вот как этого добиться:
// Благодаря мосту мы можем использовать канал каналов из одного оператора диапазона и сосредоточиться на логике нашего цикла.
// Разрушение канала каналов оставлено на усмотрение кода, специфичного для этой проблемы.
func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{}) // 1 Возвращает все значения из моста
		go func() {
			defer close(valStream)
			for { // 2 Этот цикл отвечает за извлечение каналов из chanStream и предоставление их для использования во вложенном цикле.
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) { // 3 Этот цикл отвечает за считывание значений из предоставленного ему канала и повторение этих значений в valStream
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}

}
