package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//testGoroutines1()
	//testGoroutines2()
	//testGoroutines3()
	//testGoroutines4()
	//testGoroutines5()
	testGoroutines6()

}

func testGoroutines1() {
	var ch chan int // Тут не инициирован канал
	// правильно нужно так:
	// ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}

// Что мы увидим ?
// processed: cmd.1
// processed: cmd.2
//
func testGoroutines2() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2"
}

//  Тут передача через копирование
func testGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(3 * time.Second)
}

func testGoroutines4() {
	var num int

	for i := 0; i < 10000; i++ {
		go func(i int) {
			num = i
		}(i)
	}
	fmt.Printf("NUM is %d", num) // Может быть произвольной,
	// так как горутины выполняются паралельно
}

func testGoroutines5() {
	dataMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {

			// maps не потокобезопасная, тут будет ошибка
			// fatal error: concurrent map writes
			d[fmt.Sprintf("%d", num)] = num

		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}

// Выделит один виртуальный процессор
// Но планировщик может переключать горутины при простое, поэтому тут произвольный вывод
func testGoroutines6() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
		}
	}(&x)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("x = %d.\n", x)
}
