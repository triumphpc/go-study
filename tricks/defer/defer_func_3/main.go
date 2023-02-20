package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	res := f()

	fmt.Println(res) // 3

	var cacheV []*View

	fmt.Println(cacheV == nil) // true

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
		fmt.Println(a) // 2

		a = 3 // Область видимости функции, выставляет в 3

	}()

	a = 2

	return a

}
