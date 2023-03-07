package main

import "fmt"

func main() {

	test := make(map[string][]string, 2)

	fmt.Println(test["tt"] == nil)

}
