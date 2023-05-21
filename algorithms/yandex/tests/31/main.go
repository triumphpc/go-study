package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Сортировка слиянием
// Гоше дали задание написать красивую сортировку слиянием. Поэтому Гоше обязательно надо реализовать отдельно функцию merge и функцию merge_sort.
// Функция merge принимает два отсортированных массива, сливает их в один отсортированный массив и возвращает его.
// Если требуемая сигнатура имеет вид merge(array, left, mid, right),
// то первый массив задаётся полуинтервалом [left,mid) массива array,
// а второй – полуинтервалом [mid,right) массива array.
//
// Функция merge_sort принимает некоторый подмассив, который нужно отсортировать.
// Подмассив задаётся полуинтервалом — его началом и концом.
// Функция должна отсортировать передаваемый в неё подмассив, она ничего не возвращает.
// Функция merge_sort разбивает полуинтервал на две половинки и рекурсивно вызывает сортировку отдельно для каждой.
// Затем два отсортированных массива сливаются в один с помощью merge.
// Заметьте, что в функции передаются именно полуинтервалы [begin,end), то есть правый конец не включается.
// Например, если вызвать merge_sort(arr, 0, 4), где arr=[4,5,3,0,1,2], то будут отсортированы только первые четыре элемента,
// изменённый массив будет выглядеть как arr=[0,3,4,5,1,2].

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

	scanner.Scan()
	left, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	right, _ := strconv.Atoi(scanner.Text())

	result := mergeRun(digits, left, right)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func mergeRun(data []int, left, right int) []int {
	leftPart := data[:left]
	rightPart := data[right:]
	needSort := data[left:right]
	sorted := mergeSort(needSort)

	result := make([]int, 0, len(data))
	result = append(result, leftPart...)
	result = append(result, sorted...)
	result = append(result, rightPart...)

	return result
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = mergeSort(data[:mid])
		done <- true
	}()

	right := mergeSort(data[mid:])

	<-done

	return merge(left, right)

}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged

}
