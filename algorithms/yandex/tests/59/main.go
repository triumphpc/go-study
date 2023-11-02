package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Гоша едет в гости к друзьям. Ему придётся сначала ехать на метро, а потом пересаживаться на автобус. Гоша не любит долго ждать, поэтому хочет выбрать такую станцию метро, рядом с которой расположено как можно больше остановок автобуса. Гоша считает, что остановка находится рядом с метро, если расстояние между ними не превосходит 20 метров.
//Гоше известны все координаты автобусных остановок и координаты выходов из метро. Помогите ему найти выход из метро, рядом с которым расположено больше всего остановок.
//Напомним, что расстояние между двумя точками с координатами x1, y1 и x2, y2 вычисляется по формуле .
// Пояснение к примеру 1:
//Рядом с 1-м выходом (-1, 0) находится только остановка с координатами (10, 0).
//Рядом со 2-м выходом (1, 0) находятся уже две остановки —– (10, 0) и (20, 0).
//Рядом с 3-м выходом (2, 5) расположены все три остановки –— (22, 5), (20, 0) и (10, 0)
//Пояснение к примеру 2:
//Третий выход теперь находится в точке (0, 5). Он рядом только с двумя остановками — (20, 5) и (10, 0).
//Рядом с 2-м и 3-м выходом одинаковое число остановок, поэтому в ответ идет 2-й выход, так как он раньше во входных данных.
//Формат ввода
//
//В первой строке дано количество выходов из метро –— натуральное число n (1 ≤ n ≤ 104). В следующих n строках даны координаты выходов из метро. Каждый выход описывается двумя координатами x и y, записанными через пробел.
//В следующей строке дано количество автобусных остановок —– натуральное число m (1 ≤ m ≤ 106). В следующих m строках заданы координаты остановок. Каждая остановка описывается двумя числами —– своими x и y координатами, записанными через пробел.

const maxDistance = 20

func getDistance(x1, x2, y1, y2 int) float64 {
	return math.Sqrt(math.Abs(math.Pow(float64(x1-x2), 2)) + math.Abs(math.Pow(float64(y1-y2), 2)))
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	exits := make([][]int, 0, 10)

	for i := 0; i < n; i++ {
		scanner.Scan()
		chunk := strings.Split(scanner.Text(), " ")

		ch := make([]int, len(chunk))

		for j := range chunk {
			v, _ := strconv.Atoi(chunk[j])
			ch[j] = v
		}

		exits = append(exits, ch)
	}

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	stops := make([][]int, 0, 10)

	for i := 0; i < n; i++ {
		scanner.Scan()
		chunk := strings.Split(scanner.Text(), " ")

		ch := make([]int, len(chunk))

		for j := range chunk {
			v, _ := strconv.Atoi(chunk[j])
			ch[j] = v
		}

		stops = append(stops, ch)
	}

	// Считаем расстояния
	counts := make(map[int]uint8, len(exits))
	var max uint8
	var result uint8

	for i := range exits {
		for j := range stops {
			dis := getDistance(exits[i][0], stops[j][0], exits[i][1], stops[j][1])
			if dis <= maxDistance {
				counts[i]++
				if counts[i] > max {
					max = counts[i]
					result = uint8(i) + 1
				}
			}
		}
	}

	fmt.Fprintln(dst, result)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		writer.Flush()
	}()
	task(os.Stdin, writer)
}
