package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Пузырьковая сортировка #1
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

func sort(array []int8) {
	for {
		has := false
		for i := 1; i < len(array); i++ {
			if array[i] < array[i-1] { // 1. Если следующий меньше предыдущего
				has = true
				array[i], array[i-1] = array[i-1], array[i] // Меняем местами элементы
			}
		}

		if !has {
			break
		}
	}
}
