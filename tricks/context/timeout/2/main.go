package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	ctxChild, cancel2 := context.WithCancel(ctx)

	cancel()

	fmt.Println(ctxChild.Err())

	cancel2()

	fmt.Println(ctxChild.Err())
	fmt.Println(context.Background().Err())

}
