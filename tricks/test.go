package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string
}

func changeName(person *Person) {
	//person = &Person{ // Тут переопределяется аругумент, он не является уже курсором
	//	Name: "Alice",
	//}

	// 1. Как можно изменить не меняя структуры
	//person.Name = "Alice"

	// 2. Способ
	//*person = Person{
	//	Name: "Alice",
	//}
}

type POADataJSONResponseT struct {
	Cost float64 `json:"cost,omitempty"`
}

func main() {
	res := &POADataJSONResponseT{
		Cost: 234,
	}

	data, _ := json.Marshal(res)

	log.Println(string(data))

	sss := "{\"cost\":234.43}\n"

	res2 := new(POADataJSONResponseT)

	json.Unmarshal([]byte(sss), res2)
	log.Println(int(res2.Cost))

}
