package main

// Дан максимальный массив натуральных чисел, числа могут повторяться
// Необходимо выбрать из них k чисел, чтобы разность максимального и минимального из выбранных быламинимальной
// Вернуть эту разность
// 10, 100,300, 200, 1000,20, 30 k=3 -> 20

// 1. Сначала надо отсортировать список
// 2. Далее смещаем левый маркер на один и вычисляем разницу между самым левым и самым правым
// 3. Ну и понимаем какой меньше всего

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
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

	min := parts[size-1]
	slices.Sort(parts)

	fmt.Println(parts)

	for i := 0; i < size-k; i++ {
		if parts[i+k-1]-parts[i] < min {
			min = parts[i+k-1] - parts[i]
		}
	}

	fmt.Println(min)
}
