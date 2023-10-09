package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Периметр треугольника
//Перед сном Рита решила поиграть в игру на телефоне. Дан массив целых чисел, в котором каждый элемент обозначает длину стороны треугольника.
//Нужно определить максимально возможный периметр треугольника, составленного из сторон с длинами из заданного массива. Помогите Рите скорее закончить игру и пойти спать.
//Напомним, что из трёх отрезков с длинами a ≤ b ≤ c можно составить треугольник, если выполнено неравенство треугольника: c < a + b
//
//Разберём пример:
//
//даны длины сторон 6, 3, 3, 2. Попробуем в качестве наибольшей стороны выбрать 6. Неравенство треугольника не может выполниться,
//так как остались 3, 3, 2 —– максимальная сумма из них равна 6.
//
//Без шестёрки оставшиеся три отрезка уже образуют треугольник со сторонами 3, 3, 2. Неравенство выполняется: 3 < 3 + 2. Периметр равен 3 + 3 + 2 = 8.

func main() {
	task2(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) { // Неоптимальный
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	parts := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	quickSort(parts, 0, size-1) // Используем быструю сортировку

	fmt.Fprintf(writer, "%s ", fmt.Sprint(find(parts))) // Ищем стороны
}

func find(lines []int) (res int) {
	//for idx := range arr {
	//for j := idx + 1; j < len(arr)-1; j++ {
	//	if arr[idx] < arr[j]+arr[j+1] {
	//		cur := arr[idx] + arr[j] + arr[j+1]
	//		if cur > res {
	//			res = cur
	//		}
	//	}
	//}
	//}

	for i := 0; i < len(lines)-2; i++ { // Тут O(n) поиск сторон
		if lines[i] < lines[i+1]+lines[i+2] {
			return lines[i] + lines[i+1] + lines[i+2]
		}
	}
	return res
}

func quickSort(arr []int, left, right int) {
	if right-left <= 1 {
		if right < 0 {
			return
		}
		if arr[right] < arr[left] {
			arr[left], arr[right] = arr[right], arr[left]
		}

		return
	}

	mid := left + (right-left)/2
	leftPt := left
	rightPt := right

	for leftPt < rightPt {
		if arr[leftPt] < arr[mid] {
			leftPt++
			continue
		}

		if arr[mid] < arr[rightPt] {
			rightPt--
			continue
		}

		// swap
		arr[leftPt], arr[rightPt] = arr[rightPt], arr[leftPt]
		leftPt++
		rightPt--
	}

	quickSort(arr, left, mid)
	quickSort(arr, mid, right)
}

func mergeSortDesc(arr []int) []int { // Сортировка слиянием
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSortDesc(arr[:len(arr)/2])
	right := mergeSortDesc(arr[len(arr)/2:])
	result := make([]int, 0, len(arr))

	l, r := 0, 0
	for l < len(left) && r < len(right) { // Сортировка O(n) переданных частей массива
		if left[l] > right[r] {
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

func task2(src io.Reader, dst io.Writer) { // Оптимальный
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	lines := make([]int, 0)
	for i := 0; i < count; i++ {
		scanner.Scan()
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	result := 0
	lines = mergeSortDesc(lines)

	for i := 0; i < count-2; i++ { // А тут уже просто высчитываем
		if lines[i] < lines[i+1]+lines[i+2] {
			result = lines[i] + lines[i+1] + lines[i+2]
			break
		}
	}
	fmt.Fprintf(dst, "%d\n", result)
}
