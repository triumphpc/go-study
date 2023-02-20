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
	fmt.Println("b1 == b2:", reflect.DeepEqual(b1, b2)) // b1 == b2: false
	fmt.Println("b3 == b2:", reflect.DeepEqual(b3, b2)) // b3 == b2: true

	var b4 []byte = nil
	b5 := []byte{}
	fmt.Println("b4 == b5:", bytes.Equal(b4, b5))       // prints: b4 == b5: true
	fmt.Println("b4 == b5:", reflect.DeepEqual(b4, b5)) // prints: b4 == b5: false
	fmt.Println("b4 == b1:", reflect.DeepEqual(b4, b1)) // prints: b4 == b1: true

	var str string = "one"
	var in interface{} = "one"
	fmt.Println("str == in:", str == in, reflect.DeepEqual(str, in)) // str == in: true true

	v1 := []string{"one", "two"}
	v2 := []interface{}{"one", "two"}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2)) // v1 == v2: false

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}

	data2 := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}

	fmt.Printf("data == data2: %t, %v, %v, %T, %T\n", reflect.DeepEqual(data, data2), data, data2, data, data2)
	// data == data2: true, map[code:200 value:[one two]], map[code:200 value:[one two]], map[string]interface {}, map[string]interface {}

	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)

	fmt.Printf("data == decoded: %t, %#v, %#v, %T, %T\n", reflect.DeepEqual(data, decoded), data, decoded, data, decoded)
	// data == decoded: false, map[string]interface {}{"code":200, "value":[]string{"one", "two"}}, map[string]interface {}{"code":200, "value":[]interface {}{"one", "two"}}, map[string]interface {}, map[string]interface {}

	decoded["code"] = 200
	decoded["value"] = []string{"one", "two"}
	fmt.Printf("data == decoded: %t, %v, %v, %T, %T\n", reflect.DeepEqual(data, decoded), data, decoded, data, decoded)
	//data == decoded: true, map[code:200 value:[one two]], map[code:200 value:[one two]], map[string]interface {}, map[string]interface {}

}
