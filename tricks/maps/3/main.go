package main

import "fmt"

var m = map[string]bool{
	"test":  false,
	"test2": false,
	"test3": false,
}

var parentEvtIDs = map[string]bool{
	"a5bfd710d74844cf9927b3b1b0d5f1c1": false,
}

func main() {

	defer func() {
		for evtID, has := range m {
			if !has {
				fmt.Println(evtID) // test2
			}
		}
	}()

	a := []string{"test", "test3"}

	for _, val := range a {

		if _, ok := m[val]; ok {
			m[val] = true
		}
	}

	fmt.Println(m) // map[test:true test2:false test3:true]

	var extraData map[string]string
	fmt.Println(len(extraData)) // 0

	fmt.Println(extraData == nil) // true

	extraData = make(map[string]string)
	fmt.Println(extraData == nil) // false

	extraData["test"] = "vvv"
	fmt.Printf("%v", extraData)

	has := make(map[string]bool)
	has["one"] = true

	if _, ok := parentEvtIDs["a5bfd710d74844cf9927b3b1b0d5f1c1"]; !ok {
		println("НЕТ")
	} else {
		println("ЕСТЬ") // ЕСТЬ
	}

}
