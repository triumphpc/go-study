//https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/9780135261651-UGP2_01_06_05/
package main

import (
	"log"
)

type errorString struct {
	s string
}

// It's correct implementation
// it's compare errors address
//func (c *errorString) Error() string {
//	return "It's bug"
//}
//
//func New(text string) error  {
//	return &errorString{text}
//
//}

// it's compare values of strings
func (e errorString) Error() string {
	return e.s
}

func New(text string) error {
	return errorString{text}
}

var myErr = New("bad request")

func main() {
	// if we compare value symantic, we equal strings of errors
	// and in value symantic it was true
	if err := webCall(); err == myErr {
		log.Fatalln("What?", err)
	}

	log.Println("No error")
}

func webCall() error {
	return New("bad request")
}
