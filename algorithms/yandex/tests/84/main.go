package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// Условие: дается целое число, нужно переконвертировать данное число в сумму, слагаемых должно быть не менее двух. А само разложение должно осуществляться таким образом, что произведение слагаемых будет максимальным.

//В результате необходимо получить это самое произведение.

//Пример:

//Ввод: n = 2
//
//Вывод: 1
//
//Объяснение: 2 = 1 + 1, 1 × 1 = 1.
//
//
//
//Ввод: n = 10
//
//Вывод: 36
//
//Объяснение: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

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
	num, _ := strconv.Atoi(scanner.Text())

	fmt.Println(integerBreak(num))

}

func integerBreak(ni int) int {
	var (
		maxPos float64 = 0
		i      float64 = 1
		j      float64 = 1
		n      float64 = float64(ni)
	)
	for ; n/i != 1; i++ { // Пока делится на шаг итерации (10/1)

		for ; j*i < n; j++ { // Проверяем умножение 1 * 1 < 10
			val := n - j*i                // 10 - 1 * 1 = 9
			multi := math.Pow(i, j) * val // 1 ^ 1 * 9
			if multi > maxPos {           // 9
				maxPos = multi
			}
		}
		j = 1 // Тут сбрасываем
	}
	return int(maxPos)
}
