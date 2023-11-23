package main

// 2-SUM. Поиск бинарным поиском

func main() {

}
func TwoSum(k int, in []int) []int {
	for i := 0; i < len(in); i++ { // Проходим по всем элемента массива
		numberToFind := k - in[i] // Получаем число, которое ищем
		l := i + 1                // Определяем границы поиска
		r := len(in) - 1

		for l <= r { // Пока указатели не соединятся
			mid := l + (r-l)/2           // Определяем середину отрезка
			if in[mid] == numberToFind { // Если значение середины равно искомому
				return []int{in[i], in[mid]} // Значит нашли разницу
			}

			// Если число меньше
			if numberToFind < in[mid] {
				r = mid - 1 // изменяем правую границу
			} else {
				l = mid + 1
			}
		}
	}

	return []int{}
}
