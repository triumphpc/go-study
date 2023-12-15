package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Сортировка вставкой #1
// Нативный алгоритм

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
	nums := scanner.Text()

	stringSlice := strings.Split(nums, " ")
	in := make([]int8, len(stringSlice))

	for idx := range stringSlice {
		n, _ := strconv.Atoi(stringSlice[idx])
		in[idx] = int8(n)
	}

	sort(in)

	fmt.Fprintf(dst, "%s ", fmt.Sprint(in))

}

// 5 4 2 1
// 4 5 2 1
// 2 >
// j = 2
// 4 2 5 1
// 2 4 5 1

func sort(array []int8) {
	for i := 1; i < len(array); i++ {
		insertItem := array[i] // 1. Определяем элемент для вставки
		j := i                 // 2. Помещаем текущу позицию во временную переменную

		for j > 0 && insertItem < array[j-1] { // 3. Если это не нулевой элемент он больше чем предыдущий
			array[j] = array[j-1] // меняем его местами
			j -= 1                // смещаем на шаг назад
		}

		array[0] = insertItem // 4. Выставляем текущей позиции значение
	}
}
