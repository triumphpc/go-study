package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// На IT-конференции присутствовали студенты из разных вузов со всей страны. Для каждого студента известен ID университета, в котором он учится.
//Тимофей предложил Рите выяснить, из каких k вузов на конференцию пришло больше всего учащихся.
// В первой строке дано количество студентов в списке —– n (1 ≤ n ≤ 15 000).
// Во второй строке через пробел записаны n целых чисел —– ID вуза каждого студента. Каждое из чисел находится в диапазоне от 0 до 10 000.
// В третьей строке записано одно число k.
// Выведите через пробел k ID вузов с максимальным числом участников. Они должны быть отсортированы по убыванию
// популярности (по количеству гостей от конкретного вуза). Если более одного вуза имеет одно и то же количество учащихся, то выводить их ID нужно в порядке возрастания.

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

	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Fprintf(writer, "%s ", fmt.Sprint(sort(parts, id)))
}

func sort(arr []int, need int) []int {
	cache := make(map[int]uint8, len(arr))

	for idx := range arr {
		cache[arr[idx]]++
	}

	result := make([]uint8, len(cache))
	for val, cnt := range cache {
		if cnt < cache[need] {
			continue
		}

		result[val] = 1
	}

	res := make([]int, 0, len(result))

	for idx := range result {
		if result[idx] == 0 {
			continue
		}

		res = append(res, idx)
	}

	return res
}
