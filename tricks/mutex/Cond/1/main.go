package main

// b
import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type message struct {
	cond *sync.Cond
	msg  string
}

func main() {
	//for conditionTrue() == false {
	// time.Sleep(1 * time.Millisecond)
	//}

	msg := message{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// 1
	for i := 1; i <= 3; i++ {
		go func(num int) {
			for {
				msg.cond.L.Lock()
				fmt.Println("Lock 2")
				msg.cond.Wait()
				fmt.Println("Wait")
				fmt.Printf("hello, i am worker%d. text:%s\n", num, msg.msg)
				msg.cond.L.Unlock()

				fmt.Println("UnLock 2")
			}
		}(i)
	}

	// 2
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		msg.cond.L.Lock()
		fmt.Println("Lock 1")
		msg.msg = scanner.Text()
		msg.cond.L.Unlock()
		fmt.Println("UnLock 1")

		msg.cond.Broadcast()
	}

}
