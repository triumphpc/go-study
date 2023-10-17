package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

//  Сумма четверок
// У Гоши есть любимое число S. Помогите ему найти все уникальные четвёрки чисел в массиве, которые в сумме дают заданное число S.
// В первой строке дано общее количество элементов массива n (0 ≤ n ≤ 1000).
//Во второй строке дано целое число S  .
//В третьей строке задан сам массив. Каждое число является целым и не превосходит по модулю 109.
//Формат вывода
//
//В первой строке выведите количество найденных четвёрок чисел.
//В последующих строках выведите найденные четвёрки. Числа внутри одной четверки должны быть упорядочены по возрастанию. Между собой четвёрки упорядочены лексикографически.

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	numbers := make([]int, 0, 10)
	result := make([][]int, 0, 5)

	for i := 0; i < n; i++ {
		scanner.Scan()
		intVal, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, intVal)
	}

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	history := make(map[int]bool, len(numbers))

	sort.Ints(numbers) // 1. Отсортируем
	// 1 1 2 3 4 4 5 8
	// 10 - 1 -2 = 7
	// 10 - 1 - 3 = 6
	// 10 - 1 - 5 = 4
	// 10 - 3 - 4 = 5

	// -2 -1 0 0 1 2
	//

	for i1 := 0; i1 < n; i1++ { // Шагаем вперед
		for j := i1 + 1; j < n-1 && len(history) > 0; j++ {
			for j2 := j + 1; j2 < n; j2++ {
				target := k - numbers[i1] - numbers[j] - numbers[j2] // Получаем разницу между целевой и двумя правыми

				if history[target] { // Далее смотрим, был ли слева целевой - если да - то есть в наборе 3
					result = append(result, []int{target, numbers[i1], numbers[j], numbers[j2]})
				}
			}
		}
		history[numbers[i1]] = true
	}

	log.Println(result)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
