package main

import "fmt"

type myType struct {
	my int
}

type MyVal int

type EventRequest struct {
	Code      string
	Parent    string
	Consumer  []string
	Data      interface{}
	ProdQueue string
	ConsQueue []string
}

func main() {
	var i myType               // аллокация структуры
	fmt.Println(&i == nil)     // false
	fmt.Println(myType{} == i) // true

	change(&i) // передача по указателю

	fmt.Println(i)         // {1}
	fmt.Println(&i == nil) // false

	//// Проверка, что структура не пустая
	fmt.Println(myType{} == i) // false

	res := ret()            // аллокация слайса структур
	fmt.Println(res == nil) // false

	f := &myType{
		123,
	}

	fmt.Printf("%+v\n", *f) // {my:123}

	var r2 []myType
	r2 = append(r2, myType{3})

	fmt.Println(len(r2))   // 1
	fmt.Println(r2 == nil) // false

	var newSt *myType

	newSt = &myType{my: 123}

	change2(newSt)
	fmt.Println(newSt) // &{1}

	var event *EventRequest
	event = getEvent()

	fmt.Println(event) // &{cccc  [] ffff  []}

	checkParentEvent(event)

	fmt.Println(event) // &{cccc vvvv [] ffff  []}

}

func checkParentEvent(event *EventRequest) {

	event.Parent = "vvvv"
}

func getEvent() (event *EventRequest) {
	return &EventRequest{
		Code: "cccc",
		Data: "ffff",
	}
}

func change2(i *myType) {

	*i = myType{1}
}

func change(i *myType) {

	*i = myType{1}
}

func ret() []myType {

	r := make([]myType, 1)

	return r

}
