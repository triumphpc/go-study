package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

// Полиномиальный хэш. Сломай меня
// Гоша написал программу, которая сравнивает строки исключительно по их хешам. Если хеш равен, то и строки равны. Тимофей увидел это безобразие и поручил вам сломать программу Гоши, чтобы остальным неповадно было.
//В этой задаче вам надо будет лишь найти две различные строки, которые для заданной хеш-функции будут давать одинаковое значение.
//Гоша использует следующую хеш-функцию:
// h(s) = (s1a^n-1 + s2a^n-2+....sn-1a+sn)mod m
// В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.

// Перебором всех возможных вариантов высчитываем хэш.

var hashM map[uint64]string
var alphabet []string

func task(src io.Reader, dst io.Writer) {
	alphabet = make([]string, 0)

	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	hashM = make(map[uint64]string)

	ctx, cancel := context.WithCancel(context.Background())
	o, n := parse(ctx, cancel, "")
	fmt.Fprintf(dst, "%s %s", o, n)
}

func parse(ctx context.Context, cancel context.CancelFunc, result string) (string, string) {
	for _, ch := range alphabet {
		if ctx.Err() != nil {
			return "", ""
		}

		cur := result + ch
		if len(cur) > 10 {
			return "", ""
		}

		hash := hashStr(cur)

		old, ok := hashM[hash]
		if ok {
			cancel()
			return old, cur
		}

		hashM[hash] = cur

		o, n := parse(ctx, cancel, cur)
		if o != "" {
			return o, n
		}
	}

	return "", ""
}

func hashStr(str string) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		hash += uint64(str[i]) // берем числовой значение ASCII
		if i < len(str)-1 {
			hash = (hash * uint64(1000)) % uint64(123987123) // умножаем текущее значение на основание a и берем остаток по модулю
		}
	}

	return hash
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
