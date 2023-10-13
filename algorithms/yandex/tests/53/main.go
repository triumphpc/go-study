package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Полиномиальный хэш
// Алле очень понравился алгоритм вычисления полиномиального хеша. Помогите ей написать функцию, вычисляющую хеш строки s. В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.
//Полиномиальный хеш считается по формуле:
// h(s) = (s1a^n-1 + s2a^n-2+....sn-1a+sn)mod m
// В первой строке дано число a (1 ≤ a ≤ 1000) –— основание, по которому считается хеш.
//Во второй строке дано число m (1 ≤ m ≤ 109) –— модуль.
//В третьей строке дана строка s (0 ≤ |s| ≤ 106), состоящая из больших и маленьких латинских букв.

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 0), 1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	a, _ := strconv.Atoi(scanner.Text()) // основание по которому считаем хэш
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text()) // модуль
	scanner.Scan()
	str := scanner.Text() // строка

	var hash uint64
	for i := 0; i < len(str); i++ {
		hash += uint64(str[i]) // берем числовой значение ASCII
		if i < len(str)-1 {
			hash = (hash * uint64(a)) % uint64(m) // умножаем текущее значение на основание a и берем остаток по модулю
		}
	}
	fmt.Fprintf(dst, "%d\n", hash)

	fmt.Println(-3%)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
