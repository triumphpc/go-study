package main

import (
	"fmt"
)

type MyError struct {
	Err error
}

func (e *MyError) Error() string {
	return ""
}

func myFunc() *MyError {
	return nil
}

func main() {
	// nil ошибка != nil!!!
	// https://go.dev/doc/faq#nil_error
	var err error
	err = myFunc()
	fmt.Printf("%v %T", err, err) // *main.MyError

	if err != nil {
		fmt.Println("Error") // Error
	}

}
