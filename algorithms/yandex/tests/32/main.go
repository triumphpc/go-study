package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// QuickSort быстрая сортировка

func main() {
	task(os.Stdin, os.Stdout)
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

	result := quickSort(digits)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func quickSort(array []int) []int {
	if len(array) < 2 { //  базовый случай, return array # массивы с 0 или 1 элементами фактически отсортированы
		return array
	}

	pivot := array[len(array)/2]                   // 1. Берем случайный элемент
	left, center, right := partition(array, pivot) // 2. Делим на партиции по опорной точке

	result := quickSort(left) // 4. Повторно запускаем для левой и правой части
	result = append(result, center...)
	result = append(result, quickSort(right)...)

	return result
}
func partition(array []int, pivot int) (left, center, right []int) {
	for idx := range array {
		if array[idx] > pivot { // 3. Перебираем элементы и сравниваем по опорной точке
			right = append(right, array[idx])
		} else if array[idx] < pivot {
			left = append(left, array[idx])
		} else {
			center = append(center, array[idx])
		}
	}

	return left, center, right

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

	result := mergeSortDesc(digits)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func mergeSortDesc(arr []int) []int { // Сортировка слиянием
	if len(arr) <= 1 {
		return arr
	}

	var left, right []int
	done := make(chan struct{})

	go func() {
		left = mergeSortDesc(arr[:len(arr)/2])
		done <- struct{}{}
	}()

	right = mergeSortDesc(arr[len(arr)/2:])
	<-done
	close(done)

	result := make([]int, 0, len(arr))

	l, r := 0, 0
	for l < len(left) && r < len(right) { // Сортировка O(n) переданных частей массива
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	for l < len(left) { // Тут добавляем оставшиеся части
		result = append(result, left[l])
		l++
	}
	for r < len(right) {
		result = append(result, right[r])
		r++
	}
	return result
}
