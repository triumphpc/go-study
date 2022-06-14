package main

import "fmt"

func main() {
	str := "Воттак"

	fmt.Println(str)
	fmt.Println(len(str))

	s := []rune(str)
	fmt.Println(len(s))
	ns := append(s[:3], '.', '.', '.')
	fmt.Println(string(ns))
	fmt.Println(len(ns))

	newMap := make(map[string]string)
	newMap["test1"] = "v1"
	newMap["test2"] = "v2"

	fmt.Println(newMap)

	for id, v := range newMap {
		delete(newMap, id)
		newMap["new"+v] = v
	}
	fmt.Println(newMap)

}
