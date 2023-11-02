package main

import "fmt"

// Разделить входящий поток на несколько

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

	// 1. Мы захотим использовать локальные версии out1 и out2, поэтому мы затенем эти переменные.
	//2. Мы собираемся использовать одну инструкцию select, чтобы записи в out1 и out2 не блокировали друг друга. Чтобы обеспечить запись в оба, мы выполним две итерации оператора select: по одной для каждого исходящего канала.
	//3. После записи в канал мы устанавливаем для его теневой копии значение nil, чтобы дальнейшие записи блокировались, а другой канал мог продолжаться.
	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				var out1, out2 = out1, out2 // 1
				for i := 0; i < 2; i++ {    //2

					select {
					case <-done:
					case out1 <- val:
						out1 = nil // 3
					case out2 <- val:
						out2 = nil //3
					}
				}
			}
		}()
		return out1, out2
	}

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

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
