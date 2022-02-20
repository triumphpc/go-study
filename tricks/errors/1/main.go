//https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/9780135261651-UGP2_01_06_05/
package main

import (
	"log"
)

type customErrors struct{}

func (c *customErrors) Error() string {
	return "It's bug"
}

// fail returns nil values for both return types
// but *customErrors return point on memory, not nil!  must be error object
func fail() ([]byte, *customErrors) {
	return nil, nil
}

func main() {
	var err error

	if _, err = fail(); err != nil {
		log.Fatalln("What?", err)
	}

	log.Println("No error")
}
