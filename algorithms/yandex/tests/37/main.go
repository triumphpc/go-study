package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Гардероб
// Рита решила оставить у себя одежду только трёх цветов: розового, жёлтого и малинового. После того как вещи других расцветок были убраны,
//Рита захотела отсортировать свой новый гардероб по цветам. Сначала должны идти вещи розового цвета, потом —– жёлтого, и в конце —– малинового.
//Помогите Рите справиться с этой задачей.
//Примечание: попробуйте решить задачу за один проход по массиву!

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
	parts := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	sort(parts)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(parts))
}

func sort(arr []int) {
	cache := make(map[int]int, 3) // 1. Запоминаем последнюю позицию для каждого цвета
	newArr := make([]int, 0, len(arr))

	newArr = append(newArr, arr[0])
	iMax := arr[0] // Максимальное значение

	for i := 1; i < len(arr); i++ {
		if arr[i] < iMax {
			inIdx, ok := cache[arr[i]]
			if !ok {
				// Определим какой индекс стартовый
				inIdx = 0
				for valIdx, val := range newArr {
					inIdx = valIdx
					if val > arr[i] {
						break
					}
				}
			} else {
				inIdx++
			}

			// Добавляем элемент в массив
			if inIdx > len(newArr) {
				newArr = append(newArr, arr[i])
				cache[arr[i]] = len(newArr)

				continue
			}

			tmpArr := make([]int, len(newArr)) // Добаляем новый элементы
			copy(tmpArr, newArr)
			tmp := append(tmpArr[:inIdx], arr[i])
			tmp = append(tmp, newArr[inIdx:]...)
			newArr = tmp
			cache[arr[i]] = inIdx
		} else {
			cache[arr[i]] = i
			newArr = append(newArr, arr[i])

			if arr[i] > iMax {
				iMax = arr[i]
			}
		}
	}

	copy(arr, newArr)
}
