// Если JSON-значение, которое вы пытаетесь декодировать, целочисленное, есть несколько вариантов.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	//var data = []byte(`{"status": 200}`)
	//
	//var result map[string]interface{}
	//if err := json.Unmarshal(data, &result); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}

	// Преобразовать значение с плавающей запятой в целочисленный тип, который вам нужен.
	//var status = uint64(result["status"].(float64)) // хорошо
	// Использовать тип Decoder для десериализации JSON и представления JSON-чисел с помощью интерфейсного типа Number.
	//var status, _ = result["status"].(json.Number).Int64() // хорошо
	//fmt.Println("status value:", status)

	//ver3()
	ver5()
}

// Можно использовать строковое представление вашего значения Number, чтобы десериализовать его в другой числовой тип:

func ver3() {
	var data = []byte(`{"status": 200}`)

	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		fmt.Println("error:", err)
		return
	}

	var status uint64
	if err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("status value:", status)

}

// Использовать тип struct, который преобразует (maps) числовое значение в нужный вам числовой тип.
func ver4() {
	var data = []byte(`{"status": 200}`)

	var result struct {
		Status uint64 `json:"status"`
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result => %+v", result)
	//prints: result => {Status:200}
}

// Использовать struct для преобразования числового значения в тип json.RawMessage, если требуется отложить декодирование значения.
//
//Это полезно, если вы должны выполнить декодирование условного JSON-поля (conditional field) в условиях возможности изменения структуры или типа поля.

func ver5() {
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}

}
