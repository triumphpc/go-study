package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Пузырёк
// Чтобы выбрать самый лучший алгоритм для решения задачи, Гоша продолжил изучать разные сортировки. На очереди сортировка пузырьком — https://ru.wikipedia.org/wiki/Сортировка_пузырьком
//Её алгоритм следующий (сортируем по неубыванию):
//1. На каждой итерации проходим по массиву, поочередно сравнивая пары соседних элементов. Если элемент на позиции i больше элемента на позиции i + 1, меняем их местами. После первой итерации самый большой элемент всплывёт в конце массива.
//2. Проходим по массиву, выполняя указанные действия до тех пор, пока на очередной итерации не окажется, что обмены больше не нужны, то есть массив уже отсортирован.
//3. После не более чем n – 1 итераций выполнение алгоритма заканчивается, так как на каждой итерации хотя бы один элемент оказывается на правильной позиции.

5func main() {
	task2(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	parts := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	var has bool
	for i := size - 1; i > 0; i-- { // Так как в каждой итерации самое большое число в конце, то сужаем цикл
		for j := 0; j < i; j++ {
			if parts[j] > parts[j+1] {
				has = true
				parts[j], parts[j+1] = parts[j+1], parts[j]
			}
		}

		if !has {
			break
		}

		fmt.Fprintf(out, "%s \n", fmt.Sprint(parts)) // todo понять в чем разница
		has = false
	}

	fmt.Fprintf(out, "%s ", fmt.Sprint(parts))
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	intValues := make([]int, 0, total)
	for _, strVal := range strings.Split(scanner.Text(), " ") {
		intVal, _ := strconv.Atoi(strVal)
		intValues = append(intValues, intVal)
	}
	printResults := false
	for i := total - 1; i > 0; i-- { // 1. Начинаем с полного размера (правый
		for j := 0; j < i; j++ { // 2. Левый указатель получается пока меньше правого
			if intValues[j] > intValues[j+1] { // 3. Если текущий шаг больше следующего
				intValues[j], intValues[j+1] = intValues[j+1], intValues[j] // 3. Меняем местами - получается справа всегда будут самое большое число и мы следующую итерацию делаем на одно сравнение меньше
				printResults = true
			}
		}
		if i == total-1 || printResults { // Тут сдвигаем правый шаг на один вниз
			res := fmt.Sprintf("%v", intValues)
			res = strings.TrimPrefix(res, "[")
			res = strings.TrimSuffix(res, "]")
			fmt.Fprintf(dst, "%s\n", strings.TrimSpace(res))
		}
		printResults = false
	}
}
