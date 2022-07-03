package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Что необходимо добавить вместо ABC чтобы сначала вывелись все ЗАГЛАВНЫЕ буквы а затем строчные?

const N = 26

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2 * N)

	for i := 0; i < 26; i++ {
		go func(i int) {
			defer wg.Done()

			// Gosched уступает процессор, позволяя запускать другие горутины.
			// Он не приостанавливает текущую горутину, поэтому выполнение
			// возобновляется автоматически.
			runtime.Gosched()

			fmt.Printf("%c", 'a'+1)

		}(i)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+1)

		}(i)
	}
	wg.Wait()
}
