// [Golang] Что необходимо добавить вместо ABC чтобы спарсить аргумент терминала - port?
// go run main.go -port=111
package main

import (
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
}

func main() {
	flag.Parse()
	fmt.Println(port)

}
