package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Сортировка последовательности символов

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

	sort(stringSlice)

	fmt.Fprintf(dst, "%s ", fmt.Sprint(stringSlice))

}

func less(a, b string) bool {
	return a+b < b+a
}

func sort(array []string) {
	for i := 1; i < len(array); i++ {
		insertItem := array[i]
		j := i

		for j > 0 && less(array[j-1], insertItem) {
			array[j] = array[j-1]
			j -= 1
		}

		array[j] = insertItem
		//fmt.Printf("step %d, sorted {%d} elements: {%v}\n", i, i+1, array)
	}
}
