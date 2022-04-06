// Помимо невысокой скорости (что может быть критично для вашего приложения), DeepEqual() имеет свои собственные подводные камни.
// DeepEqual() не считает пустой слайс эквивалентным nil-слайсу. Это поведение отличается от того, что вы получите
//при использовании функции bytes.Equal(): она считает эквивалентными nil и пустые слайсы.
// DeepEqual() не всегда идеальна при сравнении слайсов.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var b1 []byte = nil
	b2 := []byte{}
	b3 := []byte{}
	fmt.Println("b1 == b2:", reflect.DeepEqual(b1, b2)) // prints: b1 == b2: false
	fmt.Println("b3 == b2:", reflect.DeepEqual(b3, b2)) // prints: b1 == b2: false

	var b4 []byte = nil
	b5 := []byte{}
	fmt.Println("b4 == b5:", bytes.Equal(b4, b5)) // prints: b1 == b2: true

	var str string = "one"
	var in interface{} = "one"
	fmt.Println("str == in:", str == in, reflect.DeepEqual(str, in))
	//prints: str == in: true true

	v1 := []string{"one", "two"}
	v2 := []interface{}{"one", "two"}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: false (not ok)

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}

	data2 := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}

	fmt.Println("data == data2:", reflect.DeepEqual(data, data2), data, data2)

	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)

	fmt.Println("data == decoded:", reflect.DeepEqual(data, decoded), data, decoded)
	//prints: data == decoded: false (not ok)
}
