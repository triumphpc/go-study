package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	for i := 0; i < 10; i++ {
		doRequest(ctx)
	}

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second * 2)

}

func doRequest(ctx context.Context) {
	tm := time.NewTimer(2)

	select {
	case <-time.After(time.Second * 2): // У него есть особенность NewTimer(d).C - он возвращает и не останавливает его
		fmt.Println("timer timeout")
	case <-ctx.Done():
		tm.Stop() // Так правильней!!!

		fmt.Println("cancel")
	}
}
