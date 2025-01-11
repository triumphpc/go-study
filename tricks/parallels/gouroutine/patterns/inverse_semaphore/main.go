// Как работает инверсный семафор
// Инверсный семафор может быть реализован так, чтобы он блокировал доступ к ресурсу до тех пор, пока не будет достигнуто определенное количество "разрешений". Это может быть полезно в ситуациях, когда необходимо дождаться завершения нескольких независимых процессов перед выполнением следующего шага.
package main

import (
	"fmt"
	"sync"
)

// InverseSemaphore - структура, представляющая инверсный семафор
type InverseSemaphore struct {
	counter int
	limit   int
	done    chan struct{}
	mu      sync.Mutex
}

// NewInverseSemaphore - создает новый инверсный семафор с заданным лимитом
func NewInverseSemaphore(limit int) *InverseSemaphore {
	return &InverseSemaphore{
		limit: limit,
		done:  make(chan struct{}),
	}
}

// Signal - увеличивает счетчик и открывает доступ, если счетчик достигает лимита
func (is *InverseSemaphore) Signal() {
	is.mu.Lock()
	defer is.mu.Unlock()
	is.counter++
	if is.counter >= is.limit {
		close(is.done)
	}
}

// Wait - ожидает, пока счетчик не достигнет лимита
func (is *InverseSemaphore) Wait() {
	<-is.done
}

func main() {
	// Создаем инверсный семафор с лимитом 3
	is := NewInverseSemaphore(3)

	var wg sync.WaitGroup

	// Запускаем 3 горутины, каждая из которых сигнализирует о завершении
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Task %d completed\n", id)
			is.Signal()
		}(i)
	}

	// Ожидаем, пока все сигналы не будут получены
	is.Wait()
	fmt.Println("All tasks have been completed")
	wg.Wait()
}
