package main

import (
	"errors"
	"fmt"
)

type Requisites struct {
	Name  string `json:"name,omitempty"`
	Unit  string `json:"unit,omitempty"`
	INN   string `json:"inn,omitempty"`
	OGRN  string `json:"ogrn,omitempty"`
	Email string `json:"email,omitempty"`
}

type REQ struct {
	Req *Requisites
}

func main() {
	fmt.Println(test())

}

func test() (err error) {

	defer func() {
		if err != nil {
			fmt.Printf("here")
			//err = nil
		}
	}()

	err = errors.New("Test")

	return err

}
