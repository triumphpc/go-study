// При перенарезке получившийся слайс будет ссылаться на массив исходного слайса.
//Не забывайте об этом. Иначе может возникнуть непредсказуемое потребление памяти, когда приложение разместит в
//ней крупные временные слайсы и создаст из них новые, чтобы ссылаться на небольшие куски исходных данных.

// Чтобы избежать этой ошибки, удостоверьтесь, что копируете нужные данные из временного слайса (вместо перенарезки).
package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // выводит: 10000 10000 <byte_addr_x>
	return raw[:3]
}

func getCorrect() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // выводит: 10000 10000 <byte_addr_x>
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // выводит: 3 10000 <byte_addr_x>

	data = getCorrect()
	fmt.Println(len(data), cap(data), &data[0]) // выводит: 3 3 <byte_addr_x>
}
