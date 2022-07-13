package main

func f(...interface{}) {

}
func main() {
	f(nil...)
	//f([]int{1,2,3}...) // так нельзя использовать
	f([]interface{}{1, 2, 3}...)

}
