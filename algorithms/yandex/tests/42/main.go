package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Золотая середина
// На каждом острове в архипелаге Алгосы живёт какое-то количество людей или же остров необитаем (тогда на острове живёт 0 людей).
//Пусть на i-м острове численность населения составляет ai. Тимофей захотел найти медиану среди всех значений численности населения.
//Определение: Медиана (https://ru.wikipedia.org/wiki/Медиана_(статистика)) массива чисел a_i —– это такое число,
//что половина чисел из массива не больше него, а другая половина не меньше. В общем случае медиану массива можно найти,
//отсортировав числа и взяв среднее из них. Если количество чисел чётно, то возьмём в качестве медианы полусумму соседних средних чисел, (a[n/2] + a[n/2 + 1])/2.
//У Тимофея уже есть отдельно данные по северной части архипелага и по южной, причём значения численности населения в каждой группе отсортированы по неубыванию.
//Определите медианную численность населения по всем островам Алгосов.
//Подсказка: Если n –— число островов в северной части архипелага, а m –— в южной, то ваше решение должно работать за .

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	size2, _ := strconv.Atoi(scanner.Text())
	parts := make([]int, 0, size+size2)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	for idx := 0; idx < size2; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	if len(parts) == 0 {
		return
	}

	parts = mergeSortDesc(parts)

	if (size+size2)%2 == 0 {
		fmt.Fprintf(out, "%s ", fmt.Sprint(float32(parts[(size+size2)/2]+parts[(size+size2)/2+1])/2))
	} else {
		fmt.Fprintf(out, "%s ", fmt.Sprint(float32(parts[(size+size2)/2])))
	}
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
