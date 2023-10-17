package main

// В контексте многозадачности голодание означает, что процесс не может получить необходимые ему ресурсы в течение длительного периода времени.
//В контексте мьютексов постановка более конкретно определяется как ситуация, когда процесс не может получить блокировки READ / WRITE.
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex

	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				fmt.Println("p1")
				mu.Unlock()
			}
		}
	}()

	// goroutine 2 // не может получить блокировку долгое время
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		fmt.Println("p2")
		mu.Unlock()
	}
	done <- true
}
