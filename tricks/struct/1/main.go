package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One   int    `json:"one"`
	two   string `json:"two"`
	Three *bool  `json:"three,omitempty"`
}

func main() {
	ff := false
	in := MyData{1, "two", &ff}
	fmt.Printf("%#v\n", in) // main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // {"one":1}

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // main.MyData{One:1, two:""}

	in2 := new(MyData)
	in2.One = 123
	encoded, _ = json.Marshal(in2)
	fmt.Println(string(encoded))

	var out2 MyData
	json.Unmarshal(encoded, &out2)
	fmt.Printf("%#v\n", out2)

}
