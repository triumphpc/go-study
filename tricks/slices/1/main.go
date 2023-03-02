// Appending slices
package main

import "fmt"

func main() {
	var slice1 = make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1) // x x x x x 0 0 0

	slice2 := slice1[2:4] // - - x x 0 0 0 0
	inspectSlice(slice2)  // Length:[2], Cap [6]

	slice2[0] = "CHANGED" // c
	inspectSlice(slice1)  // x x c x x 0 0 0 // Length:[5], Cap [8]
	inspectSlice(slice2)  // - - c x 0 0 0 0 // Length:[2], Cap [6]

	slice2 = append(slice2, "NEW") // n
	inspectSlice(slice1)           // x x c x n 0 0 0 // Length:[5], Cap [8]
	inspectSlice(slice2)           // - - c x n 0 0 0 // Length:[3], Cap [6]

	// Copy with capacity right without side effect
	// Here copy base array
	slice3 := slice1[2:4:4] // новый слайс с новым базовым массивом
	// Length:[2], Cap [2]
	//[0] [0xc000014410] [CHANGED]
	//[1] [0xc000014410] [Grape]
	// 2 - начало среза (включая)
	// 4 - конец среза (не включая)
	// 4 - граница cap для нового слайса

	inspectSlice(slice3)

	slice3 = append(slice3, "NEW 2") // Сделает аллокацию, так как cap уже максимален
	inspectSlice(slice1)             // x x c x n 0 0 0 // Length:[5], Cap [8] не менялась
	inspectSlice(slice3)             // Length:[3], Cap [4]
	//[0] [0xc0000144a0] [CHANGED]
	//[1] [0xc0000144a0] [Grape]
	//[2] [0xc0000144a0] [NEW 2]
	// Видим, что cap x 2

	slice4 := make([]string, len(slice1))
	copy(slice4, slice1) // копирует без side-effect

	inspectSlice(slice4)
	// Length:[5], Cap [5]
	// [0] [0xc0000144e0] [Apple]
	// [1] [0xc0000144e0] [Orange]
	// [2] [0xc0000144e0] [CHANGED]
	// [3] [0xc0000144e0] [Grape]
	// [4] [0xc0000144e0] [NEW]

	slice4[0] = "UP"
	inspectSlice(slice4)
	// Length:[5], Cap [5]
	//[0] [0xc000014540] [UP]
	//[1] [0xc000014540] [Orange]
	//[2] [0xc000014540] [CHANGED]
	//[3] [0xc000014540] [Grape]
	//[4] [0xc000014540] [NEW]

	inspectSlice(slice1) // Неизменный

}

func inspectSlice(s []string) {
	fmt.Printf("Length:[%v], ", len(s))
	fmt.Printf("Cap [%v]\n", cap(s))

	for i, v := range s {
		fmt.Printf("[%v] [%p] [%v]\n", i, &v, v)
	}
}
