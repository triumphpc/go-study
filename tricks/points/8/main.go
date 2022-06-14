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
	var i myType

	change(&i)

	//if i == nil {
	//
	//}

	fmt.Println(i)
	// Проверка, что структура не пустая
	fmt.Println(myType{} == i)

	res := ret()

	if res == nil {
		fmt.Println("vvvv")
	}

	f := &myType{
		123,
	}

	fmt.Printf("%v", *f)

	var r2 []myType
	r2 = append(r2, myType{3})

	fmt.Println(len(r2))
	fmt.Println(r2 == nil)

	var newSt *myType

	newSt = &myType{my: 123}

	change2(newSt)
	fmt.Println(newSt)

	var event *EventRequest

	event = getEvent()

	fmt.Println(event)

	checkParentEvent(event)

	fmt.Println(event)

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

	//*i = myType{1}
}

func ret() []myType {

	r := make([]myType, 1)

	return r

}
