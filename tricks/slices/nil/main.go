package main

import "fmt"

func main() {
	var list []int
	fmt.Println(list == nil)    // true
	fmt.Println(len(list) == 0) // true

	list = []int{}
	println(list == nil)        // false
	fmt.Println(len(list) == 0) // true

	var a []int = nil
	a, a[0] = []int{1, 2}, 9 //panic: runtime error: index out of range [0]

	//fmt.Println(append([]string(nil), ""))
	fmt.Println(append([]string(nil), []string(nil)...))

	//Наводящие вопросы: каков будет результат append([]string(nil), "")? А append([]string(nil), []string(nil)...)? А почему? А range append([]string(nil), []string(nil)...) как отработает?

}
