package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Комбинации
// На клавиатуре старых мобильных телефонов каждой цифре соответствовало несколько букв. Примерно так:
//2:'abc',
//3:'def',
//4:'ghi',
//5:'jkl',
//6:'mno',
//7:'pqrs',
//8:'tuv',
//9:'wxyz'
//
//Вам известно в каком порядке были нажаты кнопки телефона, без учета повторов. Напечатайте все комбинации букв, которые можно набрать такой последовательностью нажатий.
// На вход подается строка, состоящая из цифр 2-9 включительно. Длина строки не превосходит 10 символов.

func main() {
	task2(os.Stdin, os.Stdout)

}

var mapData = map[int8][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

var in []int8

func task(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	num := scanner.Text()

	stringSlice := strings.Split(num, "")
	in = make([]int8, len(stringSlice))
	for idx := range stringSlice {
		n, _ := strconv.Atoi(stringSlice[idx])
		in[idx] = int8(n)
	}

	if len(in) == 0 {
		return
	}

	if len(in) == 1 {
		n, _ := strconv.Atoi(num)
		fmt.Fprintf(dst, "%s\n", mapData[int8(n)])
		return
	}

	parse("", in[0], 0, dst)

}

func parse(result string, pos int8, idx int, dst io.Writer) {
	for i := range mapData[pos] {
		res := result + mapData[pos][i]

		if len(in)-1 == idx {
			fmt.Fprintf(dst, "%s\n", res)

			continue
		}

		parse(res, in[idx+1], idx+1, dst)
	}
}

var keys = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	inputKeys := strings.Split(scanner.Text(), "")
	keyValues := make([][]string, 0)

	for _, key := range inputKeys {
		keyValues = append(keyValues, strings.Split(keys[key], "")) // 1. Получаем массив массивов комбинаций
	}

	if len(keyValues) == 0 {
		return
	}

	out := ""
	combine("", keyValues[1:], &out)
	fmt.Fprint(dst, out+"\n")
}

func combine(prefix string, keyValues [][]string, out *string) {
	if len(keyValues) == 0 {
		*out += fmt.Sprintf("%s ", prefix)

		return
	}

	for _, letter := range keyValues[0] { // 2. Пока не закончится последовательность, набиваем строку и потом распечатываем (можно на самом деле писать в буфер не печатая
		combine(prefix+letter, keyValues[1:], out)
	}
}
