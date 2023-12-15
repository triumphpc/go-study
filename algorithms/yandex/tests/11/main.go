package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Сортировка слиянием
// mergeSort - С использованием каналов
// mergeSort2 - наивный (базовый)

func main() {
	task(os.Stdin, os.Stdout)
	//task2(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	size, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	digits := make([]int, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()

		num, _ := strconv.Atoi(scanner.Text())
		digits[idx] = num
	}

	result := mergeSort(digits)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))

}

func task2(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	size, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	digits := make([]int, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()

		num, _ := strconv.Atoi(scanner.Text())
		digits[idx] = num
	}

	result := mergeSort2(digits)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))

}

func mergeSort2(data []int) []int {
	if len(data) <= 1 { // базовый случай рекурсии
		return data
	}

	// Запуск на левой стороне
	left := mergeSort2(data[:len(data)/2])

	// На правой
	right := mergeSort2(data[len(data)/2:])

	// Аллокация результата
	result := make([]int, len(data))

	// Слияние результатов
	l, r, k := 0, 0, 0

	for l < len(left) && r < len(right) {
		// Выбираем из какого массива забрать минимальный элемента
		if left[l] < right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}

	// Если один массив закончился раньше, чем второй то
	// переносим оставшиеся элементы второго массива в результирующие

	for l < len(left) {
		result[k] = left[l]
		l++
		k++
	}

	for r < len(right) {
		result[k] = right[r]
		r++
		k++
	}

	return result

}

// Сортировка слиянием слайсами
func mergeSort(data []int) []int {
	if len(data) <= 1 { // Если остался один элемент - его и возвращаем
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2 // 1. Сначала на две части делим
	var left []int

	go func() {
		left = mergeSort(data[:mid]) // 2. Отправляем левую часть
		done <- true
	}()

	right := mergeSort(data[mid:]) // 3. Отправляем правую часть
	<-done                         // 4. Ждем пока отработает правая часть

	res := merge(left, right) // 5. мержим левую и правую часть
	fmt.Printf("left %v right %v res %v\n ", left, right, res)

	return res

}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 { // 6. Проходимся по всем переданным элементами и сортируем в merged
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] { // Тут просто смотрим какие первые элементы больше и аппендим
			merged = append(merged, left[0])
			left = left[1:] // Затем смещаем срез
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged

}
