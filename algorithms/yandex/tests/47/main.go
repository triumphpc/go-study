package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// Хэш-функции. Соревнования

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	bal := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		intVal, _ := strconv.Atoi(scanner.Text())
		bal += intVal
	}
	nulls := n - bal
	bal = bal - nulls
	bal = n - int(math.Abs(float64(bal)))

	fmt.Fprintf(dst, "%d\n", bal)
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 0), 2*1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	resToPos := make(map[int][]int)
	prev := 0
	maxLen := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		res := scanner.Text()
		if res == "0" {
			prev = prev - 1
		}
		if res == "1" {
			prev = prev + 1
		}
		if prev == 0 && maxLen < i+1 {
			maxLen = i + 1
		}
		if len(resToPos[prev]) == 0 {
			resToPos[prev] = append(resToPos[prev], i)
		} else {
			if i-resToPos[prev][0] > maxLen {
				maxLen = i - resToPos[prev][0]
			}
		}
	}
	fmt.Fprintf(dst, "%d\n", maxLen)
}
