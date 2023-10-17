package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Общий подмассив
// Гоша увлёкся хоккеем и часто смотрит трансляции матчей. Чтобы более-менее разумно оценивать силы команд, он сравнивает очки, набранные во всех матчах каждой командой.
//Гоша попросил вас написать программу, которая по результатам игр двух выбранных команд найдёт наибольший по длине отрезок матчей, когда эти команды зарабатывали одинаковые очки.
// Рассмотрим первый пример:
//Результаты первой команды: [1 2 3 2 1].
//Результаты второй команды: [3 2 1 5 6].
//Наиболее продолжительный общий отрезок этих массивов имеет длину 3 –— это [3 2 1].
//Формат ввода
//
//В первой строке находится число n (1 ≤ n ≤ 10000) –— количество матчей, которые были сыграны первой командой.
//Во второй строке записано n целых чисел –— очки в этих играх.
//В третьей строке дано число m (1 ≤ m ≤ 10000) —– количество матчей, которые сыграла вторая команда.
//В четвертой строке заданы m целых чисел —– результаты второй команды.
//Число очков, заработанных в одной игре, лежит в диапазоне от 0 до 255.

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size1, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	nums1 := make([]int, 0, size1)
	for idx := 0; idx < size1; idx++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		nums1 = append(nums1, a)
	}

	scanner.Scan()
	size2, _ := strconv.Atoi(scanner.Text())
	nums2 := make([]int, 0, size2)
	for idx := 0; idx < size2; idx++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		nums2 = append(nums2, a)
	}

	hasNumIdx := make(map[int]int, size1)
	maxQueue := 0

	// 1. Проходимся по первому и ищем во втором
	for k, n1 := range nums1 {
		startIdx, _ := hasNumIdx[n1]
		for i := startIdx; i < len(nums2); i++ {
			if nums2[i] == n1 {

				currentQueue := 1
				// Если еще не был сдвинут флаг - проставим его
				if startIdx == 0 {
					hasNumIdx[n1] = i
				}

				// Далее вычислим максимальную длинну текущей последовательности
				j2 := i + 1

				for j := k + 1; j < len(nums1) && j2 < len(nums2); j++ {
					if nums1[j] == nums2[j2] {
						currentQueue++
						j2++
					} else {
						break
					}
				}

				if currentQueue > maxQueue {
					maxQueue = currentQueue
				}
			}
		}
	}

	fmt.Fprint(out, maxQueue)

}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
