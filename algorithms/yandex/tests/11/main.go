package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Сортировка слиянием

func main() {
	task(os.Stdin, os.Stdout)
}

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

	digits := make([]int, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()

		num, _ := strconv.Atoi(scanner.Text())
		digits[idx] = num
	}

	result := mergeSort(digits)

	outStr := ""
	for idx := range result {
		outStr += strconv.Itoa(result[idx]) + " "
	}

	_, _ = writer.Write([]byte(outStr))
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = mergeSort(data[:mid])
		done <- true
	}()

	right := mergeSort(data[mid:])

	<-done

	return merge(left, right)

}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged

}
