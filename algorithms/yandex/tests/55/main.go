package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Префиксный хэш.
// Алла не остановилась на достигнутом –— теперь она хочет научиться быстро вычислять хеши произвольных подстрок данной строки. Помогите ей!
//На вход поступают запросы на подсчёт хешей разных подстрок. Ответ на каждый запрос должен выполняться за O(1). Допустимо в начале работы программы сделать предподсчёт для дальнейшей работы со строкой.
//Напомним, что полиномиальный хеш считается по формуле
// В первой строке дано число a (1 ≤ a ≤ 1000) –— основание, по которому считается хеш. Во второй строке дано число m (1 ≤ m ≤ 107) –— модуль. В третьей строке дана строка s (0 ≤ |s| ≤ 106), состоящая из больших и маленьких латинских букв.
// В четвертой строке дано число запросов t –— натуральное число от 1 до 105. В каждой из следующих t строк записаны через пробел два числа l и r –— индексы начала и конца очередной подстроки. (1 ≤ l ≤ r ≤ |s|).

// Для каждого запроса выведите на отдельной строке хеш заданной в запросе подстроки.

func mod(n, mod int) int {
	return n % mod
}

// Возвдение в степень бинарное
func binpowFast(a, n, m int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res *= a
		}
		a *= a
		n >>= 1
	}

	return res % m
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 0), 1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	str := scanner.Text()
	hash := 0
	prefixHashes := make([]int, len(str), len(str))

	for i := 0; i < len(str); i++ {
		hash = mod(hash*int(a)+int(str[i]), m)
		prefixHashes[i] = hash
	}

	fmt.Println(prefixHashes)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())
		if l == 1 {
			fmt.Fprintf(dst, "%d\n", prefixHashes[r-1])
			continue
		}
		hashPow := r - l + 1
		left := prefixHashes[r-1]
		//right := prefixHashes[l-2] * binpowFast(a, hashPow, m)
		right := prefixHashes[l-2] * binpowFast(a, hashPow, m)
		fmt.Fprintf(dst, "%d\n", mod(left-right, m))
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
