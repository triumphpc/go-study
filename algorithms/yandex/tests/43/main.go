package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Разность треш-индексов
// Гоша долго путешествовал и измерил площадь каждого из n островов Алгосов, но ему этого мало! Теперь он захотел оценить, насколько разнообразными являются острова в составе архипелага.
//Для этого Гоша рассмотрел все пары островов (таких пар, напомним, n * (n-1) / 2) и посчитал попарно разницу площадей между всеми островами. Теперь он собирается упорядочить полученные разницы, чтобы взять k-ую по порядку из них.
//Помоги Гоше найти k-ю минимальную разницу между площадями эффективно.
//Пояснения к примерам
//Пример 1
//Выпишем все пары площадей и найдём соответствующие разницы
//1. |2 - 3| = 1
//2. |3 - 4| = 1
//3. |2 - 4| = 2
//Так как нам нужна 2-я по величине разница, то ответ будет 1.
//Пример 2
//У нас есть два одинаковых элемента в массиве —– две единицы, поэтому минимальная (первая) разница равна нулю

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	if size == 0 {
		return
	}
	parts := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	scanner.Scan()
	need, _ := strconv.Atoi(scanner.Text())

	// 1. Высчитываем все разницы
	next := 1
	si := 0
	first := true
	sortParts := make([]int, 0, size*3)

	for i := 0; i < len(parts)+1; {
		sortParts = append(sortParts, absInt(parts[i]-parts[next]))
		if first {
			i++
		}

		next++
		if next == len(parts) {
			if i == len(parts)-3 {
				break
			}

			first = false
			i = si
			si++
			next = i + 2
		}
	}

	// 2. Сортируем
	sortParts = sort(sortParts)

	fmt.Fprintf(out, "%s ", fmt.Sprint(sortParts[need-1]))

}

func sort(arr []int) []int { // 2. Сортировка слиянием
	if len(arr) <= 1 {
		return arr
	}

	var left, right []int
	done := make(chan struct{})

	go func() {
		left = sort(arr[:len(arr)/2])
		done <- struct{}{}
	}()

	right = sort(arr[len(arr)/2:])
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

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	if total == 0 {
		return
	}

	squares := make([]int, 0)
	for i := 0; i < total; i++ {
		scanner.Scan()
		intVal, _ := strconv.Atoi(scanner.Text())
		squares = append(squares, intVal)
	}
	squares = sort(squares)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	diffResults := make([]int, squares[len(squares)-1], squares[len(squares)-1])
	for distanceWithPrew := 1; distanceWithPrew <= total-1; distanceWithPrew++ {
		for i := distanceWithPrew; i < total; i++ {
			diffResults[squares[i]-squares[i-distanceWithPrew]]++
		}
	}
	diffCounter := 0
	for i := 0; i < len(diffResults); i++ {
		if diffResults[i] == 0 {
			continue
		}
		diffCounter += diffResults[i]
		if diffCounter >= k {
			fmt.Fprintf(dst, "%d\n", i)
			return
		}
	}
}
