package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Поиск в сломанном массиве
// Алла ошиблась при копировании из одной структуры данных в другую. Она хранила массив чисел в кольцевом буфере.
// Массив был отсортирован по возрастанию, и в нём можно было найти элемент за логарифмическое время. Алла скопировала данные
// из кольцевого буфера в обычный массив, но сдвинула данные исходной отсортированной последовательности.
// Теперь массив не является отсортированным. Тем не менее, нужно обеспечить возможность находить в нем элемент за O(logn).
// Можно предполагать, что в массиве только уникальные элементы.
// o(log n) это бинарный поиск
// 19 21 100 101 1 4 5 7 12
func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	need, _ := strconv.Atoi(scanner.Text())

	digits := make([]int, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		digits[idx] = num
	}

	result := binarySearch(digits, 0, size-1, need)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func binarySearch(arr []int, left, right, need int) int {
	// Частный случай
	if right-left <= 1 {
		if arr[left] == need {
			return left
		}

		if arr[right] == need {
			return right
		}

		return -1
	}

	mid := left + (right-left)/2 // 1. Берем средний элемент со смещением

	if arr[left] <= arr[mid] { // 2. Проверим, что левая часть упорядочена
		if need >= arr[left] && need < arr[mid] {
			return binarySearch(arr, left, mid, need)
		}

		// Иначе в неупорядочной части
		return binarySearch(arr, mid, right, need)
	} else {
		if need <= arr[right] && need > arr[mid] {
			return binarySearch(arr, mid, right, need)
		}

		return binarySearch(arr, left, mid, need)
	}
}
