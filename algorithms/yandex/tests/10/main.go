package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Ближайший ноль
// Улица, на которой хочет жить Тимофей, имеет длину n, то есть состоит из n одинаковых идущих подряд участков.
// На каждом участке либо уже построен дом, либо участок пустой. Тимофей ищет место для строительства своего дома.
// Он очень общителен и не хочет жить далеко от других людей, живущих на этой улице.
// Чтобы оптимально выбрать место для строительства, Тимофей хочет для каждого участка знать расстояние до ближайшего пустого участка.
//(Для пустого участка эта величина будет равна нулю –— расстояние до самого себя).
// Ваша задача –— помочь Тимофею посчитать искомые расстояния. Для этого у вас есть карта улицы.
// Дома в городе Тимофея нумеровались в том порядке, в котором строились, поэтому их номера на карте никак не упорядочены.
// Пустые участки обозначены нулями.

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	size, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	// Проходимся по каждому дому и вычисляем для него координаты
	houseCrd := make([]int, size)
	currentLen := 0
	leftIdx := 0
	currentLeft := 0

	for num := 0; num < size; num++ {
		scanner.Scan()

		houseNum, _ := strconv.Atoi(scanner.Text())

		currentLen++

		if houseNum == 0 {
			currentLen = 0
			if num > 1 {
				houseCrd[num-1] = 1
			}

			if leftIdx == 0 {
				currentLeft = 0
			} else {
				currentLeft = leftIdx + ((num - leftIdx) / 2)

			}

			for j := num - 2; j >= currentLeft; j-- {
				houseCrd[j] = houseCrd[j+1] + 1
			}

			leftIdx = num + 1
		}

		houseCrd[num] = currentLen
	}

	outStr := ""
	for idx := range houseCrd {
		outStr += strconv.Itoa(houseCrd[idx]) + " "
	}

	_, _ = writer.Write([]byte(outStr))
}

// Реализация Ворона
func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	writer := bufio.NewWriter(dst)

	defer func() {
		writer.WriteString("\n")
		writer.Flush()
	}()

	firstChunk := true
	chunkSize := 0

	for i := 0; i < total; i++ {
		scanner.Scan()
		num := scanner.Text()
		if num == "0" {
			if firstChunk {
				for i := chunkSize; i > 0; i-- {
					writer.WriteString(strconv.Itoa(i) + " ")
				}
				// add zero at the end of results
				writer.WriteString("0 ")
				chunkSize = 0
				firstChunk = false
				continue
			}
			middleOfChunk := chunkSize / 2
			if chunkSize%2 != 0 {
				middleOfChunk++
			}
			for i := 1; i <= middleOfChunk; i++ {
				writer.WriteString(strconv.Itoa(i) + " ")
			}
			for i := chunkSize / 2; i > 0; i-- {
				writer.WriteString(strconv.Itoa(i) + " ")
			}

			writer.WriteString("0 ")
			chunkSize = 0
			continue
		}
		chunkSize++
		if i == total-1 {
			for i := 1; i <= chunkSize; i++ {
				writer.WriteString(strconv.Itoa(i) + " ")
			}
		}
	}

}
