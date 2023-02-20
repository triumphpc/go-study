package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) // выводит 1
	{
		fmt.Println(x) // выводит 1
		//x := 2
		x = 2
		fmt.Println(x) // выводит 2
	}
	fmt.Println(x) // 1 или 2
}
