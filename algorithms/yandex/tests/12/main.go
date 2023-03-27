package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Ловкость рук
// Гоша и Тимофей нашли необычный тренажёр для скоростной печати и хотят освоить его. Тренажёр представляет собой поле из клавиш 4× 4,
// в котором на каждом раунде появляется конфигурация цифр и точек. На клавише написана либо точка, либо цифра от 1 до 9.
// В момент времени t игрок должен одновременно нажать на все клавиши, на которых написана цифра t. Гоша и Тимофей могут нажать в один момент
// времени на k клавиш каждый. Если в момент времени t были нажаты все нужные клавиши, то игроки получают 1 балл.
// Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём.

func main() {
	task2(os.Stdin, os.Stdout)
}

// Реализация Ворона
func task2(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	// firts param from input
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	// map of values
	values := make(map[string]int)
	// lines to map
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()
		for _, val := range strings.Split(line, "") {
			// skip "." values
			if val != "." {
				values[val]++
			}
		}
	}
	result := 0
	for _, count := range values {
		if count <= k*2 {
			result++
		}
	}
	writer.WriteString(strconv.Itoa(result) + "\n")
	writer.Flush()

}
