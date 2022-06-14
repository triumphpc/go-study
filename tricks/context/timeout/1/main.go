package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("for 1 sec") // for 1 sec because without for
	//case <-time.After(2 * time.Second):
	//	fmt.Println("for 2 sec")
	//case <-time.After(3 * time.Second):
	//	fmt.Println("for 3 sec")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("for 1 sec")
		case <-time.After(2 * time.Second):
			fmt.Println("for 2 sec")
		case <-time.After(3 * time.Second):
			fmt.Println("for 3 sec")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return

		}
	}
	// for 1 sec
	//for 1 sec
	//context deadline exceeded

}
