package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Дана длинная строка, состоящая из маленьких латинских букв. Нужно найти все её подстроки длины n, которые встречаются хотя бы k раз.
//Формат ввода
//
//В первой строчке через пробел записаны два натуральных числа n и k.
//Во второй строчке записана строка, состоящая из маленьких латинских букв. Длина строки 1 ≤ L ≤ 106.
//n ≤ L, k ≤ L.
//Формат вывода
//
//Для каждой найденной подстроки выведите индекс начала её первого вхождения (нумерация в строке начинается с нуля).
//Выводите индексы в любом порядке, в одну строку, через пробел.

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println(n, k)

	scanner.Scan()
	st := scanner.Text()

	subsIdx := make(map[string][]int, len(st))

	for i := 0; i < len(st)-n; i++ {
		if subsIdx[st[i:i+n]] == nil {
			subsIdx[st[i:i+n]] = make([]int, 0, 1)
		}
		subsIdx[st[i:i+n]] = append(subsIdx[st[i:i+n]], i)
	}

	// Фильтруем нужные
	for _, cn := range subsIdx {
		if len(cn) >= k {
			fmt.Fprintln(dst, cn[0])
		}
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
