package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Хэш-функции. Странное сравнение

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	//scanner.Scan()

	scanner.Scan()
	val := scanner.Text()

	scanner.Scan()
	val2 := scanner.Text()

	if len(val) != len(val2) {
		fmt.Fprintf(dst, "NO\n")

		return
	}

	nums := make(map[string]string, len(val))

	for i := 0; i < len(val); i++ {
		cur, ok := nums[string(val[i])]
		if ok {
			if cur != string(val2[i]) {
				fmt.Fprintf(dst, "NO\n")

				return
			}
		}
		nums[string(val[i])] = string(val2[i])
	}

	fmt.Fprintf(dst, "YES\n")
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 0), 1024*1024)
	scanner.Scan()
	str1 := scanner.Text()
	scanner.Scan()
	str2 := scanner.Text()
	if len(str1) != len(str2) {
		fmt.Fprint(dst, "NO")
		return
	}
	map1to2 := make(map[int]int)
	map2to1 := make(map[int]int)
	for i := 0; i < len(str1); i++ {
		code1 := int(str1[i])
		code2 := int(str2[i])
		if _, ok := map1to2[code1]; !ok {
			if _, ok := map2to1[code2]; ok {
				fmt.Fprint(dst, "NO\n")
				return
			}
			map1to2[code1] = code2
			map2to1[code2] = code1
			continue
		}
		if map1to2[code1] != code2 || map2to1[code2] != code1 {
			fmt.Fprint(dst, "NO\n")
			return
		}
	}
	fmt.Fprint(dst, "YES\n")
}
