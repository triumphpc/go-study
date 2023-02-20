// Какая функция лучший вариант для прохождения каждой строки текстового файла?

// fmt.Fscanf() - применимо только к форматированным строкам.
// bufio.Reader.ReadLine() - очень низкий уровень. Может потребоваться больше вызовов при превышении лимита буфера.
// bufio.ReadString('\n') - не может обработать EOF.
// bufio.Scanner.Scan() - лучше всего подходит.

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func main() {
	input := "foofoofoo"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if bytes.Equal(data[:3], []byte{'f', 'o', 'o'}) {
			return 3, []byte{'F'}, nil
		}
		if atEOF {
			return 0, nil, errors.New("Scan error")
		}
		return 0, nil, nil
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

}
