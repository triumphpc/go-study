package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Кружки.
// В компании, где работает Тимофей, заботятся о досуге сотрудников и устраивают различные кружки по интересам. Когда кто-то записывается на занятие, в лог вносится название кружка.
//По записям в логе составьте список всех кружков, в которые ходит хотя бы один человек.

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	circles := make(map[string]bool)
	for i := 0; i < total; i++ {
		scanner.Scan()
		circle := scanner.Text()
		if _, ok := circles[circle]; !ok {
			fmt.Fprintf(dst, "%s\n", circle)
			circles[circle] = true
		}
	}
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	circles := make(map[string]bool, total)

	for i := 0; i < total; i++ {
		scanner.Scan()
		circle := scanner.Text()
		if circles[circle] {
			fmt.Fprintf(dst, "%s\n", circle)
			circles[circle] = true
		}
	}
}
