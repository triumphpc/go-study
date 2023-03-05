package main

import (
	"encoding/json"
	"fmt"
)

type View struct {
	MessageID   string          `json:"messageId"`
	MessageType string          `json:"messageType"`
	MessageData json.RawMessage `json:"message"`
	Ver         string          `json:"ver"`
}

type DataStruct struct {
	ID   string
	Name string
}

type Pack struct {
	v View
}

func main() {

	data := DataStruct{
		ID:   "test",
		Name: "test 2",
	}

	d, _ := json.Marshal(data)

	pack := Pack{}
	pack.v.MessageType = "Вот тут "
	pack.v.MessageID = "И смотри какая тема"
	pack.v.MessageData = d

	res2, _ := json.Marshal(pack.v)

	// Теперь распакуем

	view := View{}

	json.Unmarshal(res2, &view)

	fmt.Printf("%#v %T", view, view)

	// Распаковка вложенного
	dataView := DataStruct{}
	json.Unmarshal(view.MessageData, &dataView)

	fmt.Printf("%#v %T", dataView, dataView)

}
