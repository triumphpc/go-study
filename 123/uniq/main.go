package main

import (
	"fmt"
	"unique"
)

type ss struct {
	ID   string
	Name string
}

func main() {

	pp := unique.Make(ss{ID: "fdf"})
	fmt.Println(pp)
}
