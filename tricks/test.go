package main

import (
	"fmt"
	"regexp"
)

type Requisites struct {
	Name  string `json:"name,omitempty"`
	Unit  string `json:"unit,omitempty"`
	INN   string `json:"inn,omitempty"`
	OGRN  string `json:"ogrn,omitempty"`
	Email string `json:"email,omitempty"`
}

type REQ struct {
	Req *Requisites
}

func main() {
	query := "+4794243432434"

	matched, _ := regexp.MatchString(`^(\+7|8|7).{1,100}$`, query)

	fmt.Println(matched) // false

	//fmt.Println(matched)

	//str := "{\n  \"unit\": \"ФИЛИАЛ\",\n  \"inn\": \"9616134640\",\n  \"ogrn\": \"2158233686004\",\n  \"name\": \"ООО \\\"ПОСАД-ТРАНС\\\"\",\n  \"email\": \"vrulin_sv@astralnalog.ru\",\n  \"address\": {\n    \"subject\": \"40 Калужская область\",\n    \"subRegion\": \"Калуга\",\n    \"street\": \"ул Циолковского\",\n    \"building\": \"д 10\"\n  },\n  \"owner\": {\n    \"inn\": \"966118202380\",\n    \"snils\": \"503-804-363 53\",\n    \"patronymic\": \"Тестович\",\n    \"firstname\": \"Тест\",\n    \"surname\": \"Тестов\",\n    \"post\": \"Директор\"\n  }\n}"
	////
	////srt2 := "{\"unit\": \"test\"}"
	//
	//req := new(Requisites)
	//
	//json.Unmarshal([]byte(str), req)
	//
	//fmt.Println(req.Name)
	//
	//var data json.RawMessage
	//
	//var err error
	//if data, err = json.Marshal(req); err != nil {
	//}
	//
	//fmt.Println(string(data))

}
