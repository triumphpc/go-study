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

	// 1. Новый слайс, с новой капасити, ссылается на старый базовый массив
	slice2 := slice1[2:4] // - - x x 0 0 0 0
	inspectSlice(slice2)  // Length:[2], Cap [6]

	slice2[0] = "CHANGED" // c
	inspectSlice(slice1)  // x x c x x 0 0 0 // Length:[5], Cap [8]
	inspectSlice(slice2)  // - - c x 0 0 0 0 // Length:[2], Cap [6]

	// 2. Добавляем новый элемент - меняется базовый массив
	slice2 = append(slice2, "NEW") // n
	inspectSlice(slice1)           // x x c x n 0 0 0 // Length:[5], Cap [8]
	inspectSlice(slice2)           // - - c x n 0 0 0 // Length:[3], Cap [6]

	// 3. Копирование слайса с новой капасити и с новым базовым массивом
	slice3 := slice1[2:4:4] // новый слайс с новым базовым массивом
	// Length:[2], Cap [2]
	//[0] [0xc000014410] [CHANGED]
	//[1] [0xc000014410] [Grape]
	// 2 - начало среза (включая)
	// 4 - конец среза (не включая)
	// 4 - граница cap для нового слайса

	inspectSlice(slice3) // x x

	// 4. Ростет в два раза
	slice3 = append(slice3, "NEW 2") // Сделает аллокацию, так как cap уже максимален
	inspectSlice(slice1)             // x x c x n 0 0 0 // Length:[5], Cap [8] не менялась
	inspectSlice(slice3)             // Length:[3], Cap [4]
	//[0] [0xc0000144a0] [CHANGED]
	//[1] [0xc0000144a0] [Grape]
	//[2] [0xc0000144a0] [NEW 2]
	// Видим, что cap x 2

	// 5. Копирование с новым базовым массивом
	slice4 := make([]string, len(slice1))
	copy(slice4, slice1) // копирует без side-effect - создание нового базового

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

	// 6. При инициализации с размерностью - заполняет значения по-умолчанию, при append - ростет в два раза
	slice5 := make([]string, 3)
	pointElem := &slice5[1]
	*pointElem = "2"
	slice5 = append(slice5, "1")

	inspectSlice(slice5) // Length:[4], Cap [6] // x 2 x 1 0 0
	// [0] [0xc0000965c0] []
	//[1] [0xc0000965c0] [2]
	//[2] [0xc0000965c0] []
	//[3] [0xc0000965c0] [1]

	// 7. Помним про ссылочный тип
	slice6 := slice5 // указывает на один и тот же базовый массив
	slice6[0] = "NEW VALUE"
	inspectSlice(slice5) //  Тут тоже меняет значение 0-го элемента
	inspectSlice(slice6)
	// Length:[4], Cap [6]
	//[0] [0xc0000965f0] [NEW VALUE]
	//[1] [0xc0000965f0] [2]
	//[2] [0xc0000965f0] []
	//[3] [0xc0000965f0] [1]
	//Length:[4], Cap [6]
	//[0] [0xc000096630] [NEW VALUE]
	//[1] [0xc000096630] [2]
	//[2] [0xc000096630] []
	//[3] [0xc000096630] [1]

	slice5 = append(slice5, "v5") // Добавим элемент 5 в срез slice5. Базовый массив только 4 первых символа, тут второй не поменяется
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[5], Cap [6]
	//[0] [0xc00010c670] [NEW VALUE]
	//[1] [0xc00010c670] [2]
	//[2] [0xc00010c670] []
	//[3] [0xc00010c670] [1]
	//[4] [0xc00010c670] [v5]
	//Length:[4], Cap [6]
	//[0] [0xc00010c6c0] [NEW VALUE]
	//[1] [0xc00010c6c0] [2]
	//[2] [0xc00010c6c0] []
	//[3] [0xc00010c6c0] [1]

	slice6 = append(slice6, "v5_2") // Добавим элемент 5_2 в срез slice6, и расширит базовый массив до 5 элементов
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[5], Cap [6]
	//[0] [0xc0000146b0] [NEW VALUE]
	//[1] [0xc0000146b0] [2]
	//[2] [0xc0000146b0] []
	//[3] [0xc0000146b0] [1]
	//[4] [0xc0000146b0] [v5_2]
	//Length:[5], Cap [6]
	//[0] [0xc000014700] [NEW VALUE]
	//[1] [0xc000014700] [2]
	//[2] [0xc000014700] []
	//[3] [0xc000014700] [1]
	//[4] [0xc000014700] [v5_2]

	slice6 = append(slice6, "v6")
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[5], Cap [6]
	//[0] [0xc0000987a0] [NEW VALUE]
	//[1] [0xc0000987a0] [2]
	//[2] [0xc0000987a0] []
	//[3] [0xc0000987a0] [1]
	//[4] [0xc0000987a0] [v5_2]
	//Length:[6], Cap [6]
	//[0] [0xc0000987f0] [NEW VALUE]
	//[1] [0xc0000987f0] [2]
	//[2] [0xc0000987f0] []
	//[3] [0xc0000987f0] [1]
	//[4] [0xc0000987f0] [v5_2]
	//[5] [0xc0000987f0] [v6]
	slice5 = append(slice5, "v6_1") // так же расширяет
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[6], Cap [6]
	//[0] [0xc000014890] [NEW VALUE]
	//[1] [0xc000014890] [2]
	//[2] [0xc000014890] []
	//[3] [0xc000014890] [1]
	//[4] [0xc000014890] [v5_2]
	//[5] [0xc000014890] [v6_1]
	//Length:[6], Cap [6]
	//[0] [0xc0000148f0] [NEW VALUE]
	//[1] [0xc0000148f0] [2]
	//[2] [0xc0000148f0] []
	//[3] [0xc0000148f0] [1]
	//[4] [0xc0000148f0] [v5_2]
	//[5] [0xc0000148f0] [v6_1]

	slice6 = append(slice6, "element7") // Тут уже новый базовый массив
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[6], Cap [6]
	//[0] [0xc000014940] [NEW VALUE]
	//[1] [0xc000014940] [2]
	//[2] [0xc000014940] []
	//[3] [0xc000014940] [1]
	//[4] [0xc000014940] [v5_2]
	//[5] [0xc000014940] [v6_1]
	//Length:[7], Cap [12]
	//[0] [0xc0000149a0] [NEW VALUE]
	//[1] [0xc0000149a0] [2]
	//[2] [0xc0000149a0] []
	//[3] [0xc0000149a0] [1]
	//[4] [0xc0000149a0] [v5_2]
	//[5] [0xc0000149a0] [v6_1]
	//[6] [0xc0000149a0] [element7]
	slice5 = append(slice5, "elemnt7from old") // при аллокации создается новый базовый массив и на сторый уже не ссылается
	slice5[2] = "can change?"
	inspectSlice(slice5)
	inspectSlice(slice6)
	// Length:[7], Cap [12]
	//[0] [0xc0000969e0] [NEW VALUE]
	//[1] [0xc0000969e0] [2]
	//[2] [0xc0000969e0] [can change?]
	//[3] [0xc0000969e0] [1]
	//[4] [0xc0000969e0] [v5_2]
	//[5] [0xc0000969e0] [v6_1]
	//[6] [0xc0000969e0] [elemnt7from old]
	//Length:[7], Cap [12]
	//[0] [0xc000096a60] [NEW VALUE]
	//[1] [0xc000096a60] [2]
	//[2] [0xc000096a60] []
	//[3] [0xc000096a60] [1]
	//[4] [0xc000096a60] [v5_2]
	//[5] [0xc000096a60] [v6_1]
	//[6] [0xc000096a60] [element7]

	// 8. При аллокации слайса - копируется базовый массив и новые указатели
	s7 := make([]string, 1)
	s8 := s7 // ссылочный тип, это указывает на один базовый массив
	s8[0] = "val_1"
	inspectSlice(s7)
	inspectSlice(s8)
	// Length:[1], Cap [1]
	//[0] [0xc000014b10] [val_1]
	//Length:[1], Cap [1]
	//[0] [0xc000014b30] [val_1]
	s7 = append(s7, "val_2")
	inspectSlice(s7)
	inspectSlice(s8)
	// Length:[2], Cap [2]
	//[0] [0xc00010ab10] [val_1]
	//[1] [0xc00010ab10] [val_2]
	//Length:[1], Cap [1]
	//[0] [0xc00010ab40] [val_1]
	s8 = append(s8, "val_3")
	s8[0] = "new val_1"
	inspectSlice(s7)
	inspectSlice(s8) // Уже ссылается на новый базовый массив
	//Length:[2], Cap [2]
	//[0] [0xc000096b60] [val_1]
	//[1] [0xc000096b60] [val_2]
	//Length:[2], Cap [2]
	//[0] [0xc000096b90] [new val_1]
	//[1] [0xc000096b90] [val_3]

	s9 := make([]string, 1, 2)
	s10 := s9
	s10[0] = "new_val1"
	inspectSlice(s9)
	inspectSlice(s10) // Общий базовый массив
	// Length:[1], Cap [2]
	//[0] [0xc000014c00] [new_val1]
	//Length:[1], Cap [2]
	//[0] [0xc000014c20] [new_val1]

	s9 = append(s9, "val_2")
	s9[0] = "changed val_1"
	inspectSlice(s9)  // Второй начинает на новый базовый массив
	inspectSlice(s10) // Базовый массив на элементе 1
	// Length:[2], Cap [2]
	//[0] [0xc000096c00] [changed val_1]
	//[1] [0xc000096c00] [val_2]
	//Length:[1], Cap [2]
	//[0] [0xc000096c30] [changed val_1]

	s10 = append(s10, "val_3") // Расширяется базовый массив
	inspectSlice(s9)
	inspectSlice(s10)
	// Length:[2], Cap [2]
	//[0] [0xc000188a90] [changed val_1]
	//[1] [0xc000188a90] [val_3]
	//Length:[2], Cap [2]
	//[0] [0xc000188ac0] [changed val_1]
	//[1] [0xc000188ac0] [val_3]

	s10 = append(s10, "val_4") // Аллокация - начинает ссылать на новую область
	s9[0] = "CHANGED"
	inspectSlice(s9)
	inspectSlice(s10)
	// Length:[2], Cap [2]
	//[0] [0xc00010ccb0] [CHANGED]
	//[1] [0xc00010ccb0] [val_3]
	//Length:[3], Cap [4]
	//[0] [0xc00010cce0] [changed val_1]
	//[1] [0xc00010cce0] [val_3]
	//[2] [0xc00010cce0] [val_4]

	// 9. Передача копии слайса в функцию
	s11 := make([]string, 2)
	s11[0] = "v1"
	s11[1] = "v2"
	inspectSlice(s11)
	// [0] [0xc000014d60] [v1]
	//[1] [0xc000014d60] [v2]

	chElem(s11) // Передается копия слайса с ссылкой на один базовый массив
	// Length:[2], Cap [2]
	//[0] [0xc000096d50] [v1]
	//[1] [0xc000096d50] [v2]
	inspectSlice(s11)
	// [0] [0xc000096d80] [CHANGED]
	//[1] [0xc000096d80] [v2]
	s12 := s11
	s12[0] = "RETURN"
	inspectSlice(s11)
	//[0] [0xc000096db0] [RETURN]
	//[1] [0xc000096db0] [v2]

	s11 = append(s11, "v3") // аллокация = новый базовый, новые указатели
	chElem(s11)
	inspectSlice(s11)
	// Length:[3], Cap [4]
	//[0] [0xc000014e20] [RETURN]
	//[1] [0xc000014e20] [v2]
	//[2] [0xc000014e20] [v3]
	//Length:[3], Cap [4]
	//[0] [0xc000014e60] [CHANGED]
	//[1] [0xc000014e60] [v2]
	//[2] [0xc000014e60] [v3]

	inspectSlice(s12) // а этот остается старый
	// Length:[2], Cap [2]
	//[0] [0xc00010ce60] [RETURN]
	//[1] [0xc00010ce60] [v2]

}

func chElem(sl []string) {
	inspectSlice(sl)
	sl[0] = "CHANGED"

}

func inspectSlice(s []string) {
	fmt.Printf("Length:[%v], ", len(s))
	fmt.Printf("Cap [%v]\n", cap(s))

	for i, v := range s {
		fmt.Printf("[%v] [%p] [%v]\n", i, &v, v)
	}
}
