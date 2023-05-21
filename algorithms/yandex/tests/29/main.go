package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Клумбы
//Алла захотела, чтобы у неё под окном были узкие клумбы с тюльпанам. На схеме земельного участка клумбы обозначаются
//просто горизонтальными отрезками, лежащими на одной прямой. Для ландшафтных работ было нанято n садовников.
//Каждый из них обрабатывал какой-то отрезок на схеме. Процесс был организован не очень хорошо, иногда один и тот же отрезок или
//его часть могли быть обработаны сразу несколькими садовниками. Таким образом, отрезки, обрабатываемые двумя разными садовниками,
//сливаются в один. Непрерывный обработанный отрезок затем станет клумбой. Нужно определить границы будущих клумб.
//Рассмотрим примеры.
//Пример 1:
//
//Два одинаковых отрезка [7, 8] и [7, 8] сливаются в один, но потом их накрывает отрезок [6, 10]. Таким образом, имеем две клумбы с координатами [2,3] и [6,10].
//Пример 2
//
//Отрезки [2,3], [3, 4] и [3,4] сольются в один отрезок [2,4]. Отрезок [5,6] ни с кем не объединяется, добавляем его в ответ.

// План реализации
// 1. Горутины начинают обрабатывать указнное количестов координак
// 2. Аллоцируем место под клумбы
// 3. Под сравнение делаем коппайлер, кторый возвращает 3 состояния > < =
// 4. В случае с равно компайл объединяем элемент

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	nums, _ := strconv.Atoi(scanner.Text())

	parts := make([][]int8, nums) // Выделяем память под клумбы

	// Считываем клумбы
	for i := 0; i < nums; i++ {
		scanner.Scan()
		part := scanner.Text()
		coords := strings.Split(part, " ")

		parts[i] = make([]int8, 2)

		for j := 0; j < 2; j++ {
			cc, _ := strconv.Atoi(coords[j])
			parts[i][j] = int8(cc)
		}
	}

	fmt.Println("Старт сортировки")

	// Сначала объединяем
	parts = mergeSort(parts)

	fmt.Fprintf(dst, "%s ", fmt.Sprint(parts))
}

func mergeSort(data [][]int8) [][]int8 {
	if len(data) <= 1 {
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2
	var left [][]int8

	go func() {
		left = mergeSort(data[:mid])
		done <- true
	}()

	right := mergeSort(data[mid:])

	<-done

	return merge(left, right)
}

func merge(left, right [][]int8) [][]int8 {
	merged := make([][]int8, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		}

		if len(right) == 0 {
			return append(merged, left...)
		}

		switch compare(left[0], right[0]) {
		case 2:
			right = right[1:]

			_ = merge(left, right)
		case 3:
			l := left[0][0]
			if right[0][0] < l {
				l = right[0][0]
			}

			r := left[0][1]
			if right[0][1] > r {
				r = right[0][1]
			}

			left[0] = []int8{l, r}
			right = right[1:]

			_ = merge(left, right)
		case 4:
			l := left[0][0]
			if right[0][0] < l {
				l = right[0][0]
			}

			r := left[0][1]
			if right[0][1] > r {
				r = right[0][1]
			}

			right[0] = []int8{l, r}
			left = left[1:]

			_ = merge(left, right)
		case 5:
			merged = append(merged, left[0])
			left = left[1:]
		case 6:
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged
}

func compare(x, y []int8) int8 {
	if x[0] == y[0] && x[1] == y[1] {
		return 2
	}

	if x[1]+1 > y[0] && x[1] <= y[1] {
		return 3
	}

	if y[1]+1 > x[0] && y[1] <= x[1] {
		return 4
	}

	if x[1] < y[0] {
		return 5
	}

	return 6
}
