package main

import (
	"fmt"
	"time"
)

// Ваша миссия, если вы решите принять ее, состоит в том, чтобы написать функцию, которая по двум каналам a и b
//определенного типа возвращает один канал c того же типа. Каждый элемент, полученный в a или b, будет отправлен в c, и как только a и b будут закрыты, c будет также закрыт.

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
		}

		close(c)
	}()

	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c) // Тут закрываем каналы

		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					continue
				}
				c <- v
			}
		}
	}()

	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
	time.Sleep(time.Second)
}
