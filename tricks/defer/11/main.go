package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	res := f()

	fmt.Println(res)

	var cacheV []*View

	fmt.Println(cacheV == nil)

}

type View struct {
	ID       uint
	RecordID string
	Record   json.RawMessage
	Created  time.Time
	Updated  time.Time
}

func f() (a int) {

	a = 1

	defer func() {
		fmt.Println(a)

		a = 3

	}()

	a = 2

	return a

}
