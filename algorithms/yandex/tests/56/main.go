package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Анаграммная группировка
// Вася решил избавиться от проблем с произношением и стать певцом. Он обратился за помощью к логопеду. Тот посоветовал Васе выполнять упражнение, которое называется анаграммная группировка. В качестве подготовительного этапа нужно выбрать из множества строк анаграммы.
//Анаграммы –— это строки, которые получаются друг из друга перестановкой символов. Например, строки «SILENT» и «LISTEN» являются анаграммами.
//Помогите Васе найти анаграммы.
// В первой строке записано число n —– количество строк.
//Далее в строку через пробел записаны n строк.
//n не превосходит 6000. Длина каждой строки не более 100 символов.
// Нужно вывести в отсортированном порядке индексы строк, которые являются анаграммами.
//Каждая группа индексов должна быть выведена в отдельной строке. Индексы внутри одной группы должны быть отсортированы по возрастанию. Группы между собой должны быть отсортированы по возрастанию первого индекса.
//Обратите внимание, что группа анаграмм может состоять и из одной строки. Например, если в исходном наборе нет анаграмм, то надо вывести n групп, каждая из которых состоит из одного индекса.

// 1. Сортируем
// 3. Cохраняем индексы
//

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

	res := make(map[string][]int)
	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		world := scanner.Text()

		chrs := make([]int, len(world))
		for i := range world {
			chrs = append(chrs, int(world[i]))
		}

		result := quickSort(chrs)
		var key string
		for i := range result {
			key += strconv.Itoa(result[i])
		}

		res[key] = append(res[key], idx)
	}

	for _, val := range res {
		fmt.Println(val)

	}

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

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
