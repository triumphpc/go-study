// Дан набор чисел. И дано число k. Нужно перечислить количество сумм элемента массива, составляющего в сумме число K
// O(n)сложность и O(n) памяти, потому что сохраняем в массив
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

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
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println(subArraySun(parts, k))

}

func subArraySun(nums []int, k int) int {
	answer := 0
	subArrSum := 0                          // Сумма элементов подмассивов
	prefSum := make(map[int]int, len(nums)) // Мапа для хранения сумм подмассивов
	prefSum[0] = 1

	// 1. Наичнаем пробегаться по всем элементам массива
	for i := range nums {
		subArrSum += nums[i]      // 2. Суммируем элементы значений от 0 до текущего значения
		toRemove := subArrSum - k // 3. Определим кусок, который нудно отрезать от предыдущих сумм
		has, ok := prefSum[toRemove]
		if ok {
			fmt.Println(subArrSum, toRemove)
			answer += has
		}

		prefSum[subArrSum]++ // 4. Кладем в нашу мапу количественные показатели
		fmt.Println(prefSum)
	}

	return answer
}
