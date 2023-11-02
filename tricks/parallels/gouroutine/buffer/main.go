package main

import "fmt"

func main() {
	message := make(chan string, 2)
	message <- "ff"

	fmt.Println(<-message)
	message <- "ff 3"

	fmt.Println(<-message)

}
