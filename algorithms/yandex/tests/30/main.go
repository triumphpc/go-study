package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Подпоследовательность
// Гоша любит играть в игру «Подпоследовательность»: даны 2 строки, и нужно понять, является ли первая из них
// подпоследовательностью второй. Когда строки достаточно длинные, очень трудно получить ответ на этот вопрос,
// просто посмотрев на них. Помогите Гоше написать функцию, которая решает эту задачу.
// В первой строке записана строка s.
// Во второй —- строка t.
// Обе строки состоят из маленьких латинских букв, длины строк не превосходят 150000. Строки могут быть пустыми.
// BenchmarkBasic-16     	 2537817	      1394 ns/op	    8197 B/op	       2 allocs/op
// BenchmarkBasic2-16    	 2552924	      1411 ns/op	    8197 B/op	       2 allocs/op
// PASS
func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) { // On
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	s := scanner.Text()

	scanner.Scan()
	t := scanner.Text()

	fmt.Fprintf(dst, "%t ", find(t, s))
}

func find(t, s string) bool {
	sBites := []byte(t)
	tBites := []byte(s)

	for idx := range tBites {
		if tBites[idx] == sBites[0] {
			x := idx + 1
			f := true

			for j := 1; j < len(sBites); j++ {
				if sBites[j] == tBites[x] {
					x++
					continue
				} else {
					f = false
					break
				}
			}

			if f {
				return true
			}
		}
	}

	return false
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 0), 256*1024)
	scanner.Scan()
	str1 := strings.Split(scanner.Text(), "")
	scanner.Scan()
	str2 := strings.Split(scanner.Text(), "")
	str2Pos := 0
	charsFound := 0
	for i := 0; i < len(str1); i++ {
		charToFind := str1[i]
		for str2Pos < len(str2) {
			if str2[str2Pos] == charToFind {
				charsFound++
				str2Pos++
				break
			}
			str2Pos++
		}
	}
	if charsFound == len(str1) {
		fmt.Fprint(dst, "True\n")
	} else {
		fmt.Fprint(dst, "False\n")
	}
}
