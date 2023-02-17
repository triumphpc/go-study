package main

import "fmt"

var m = map[string]bool{
	"test":  false,
	"test2": false,
	"test3": false,
}

//type Test struct {
//	extraData map[string]string
//}

var parentEvtIDs = map[string]bool{
	"a5bfd710d74844cf9927b3b1b0d5f1c1": false,
}

func main() {

	defer func() {
		for evtID, has := range m {
			if !has {
				fmt.Println(evtID)
			}
		}
	}()

	a := []string{"test", "test3"}

	for _, val := range a {

		if _, ok := m[val]; ok {
			m[val] = true
		}
	}

	var extraData map[string]string

	fmt.Println(len(extraData))

	fmt.Println(extraData == nil)

	extraData = make(map[string]string)
	extraData["test"] = "vvv"

	fmt.Printf("%v", extraData)

	has := make(map[string]bool)
	has["one"] = true

	if _, ok := parentEvtIDs["a5bfd710d74844cf9927b3b1b0d5f1c1"]; !ok {
		println("НЕТ")
	} else {
		println("ЕСТЬ")
	}

}
