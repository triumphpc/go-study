package main

import "fmt"

var m = map[string]bool{
	"test":  false,
	"test2": false,
	"test3": false,
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

}
