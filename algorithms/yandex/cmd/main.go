// https://contest.yandex.ru/contest/26365/problems/A/
package main

import (
	"flag"
	"fmt"
	"os"

	y "github.com/triumphpc/go-study/algorithms/yandex"
	tt "github.com/triumphpc/go-study/algorithms/yandex/tests/1"
)
import "bufio"
import "strings"
import "strconv"

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := make([]int64, 2)

	line1, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(line1), 10, 32)

	line2, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseInt(strings.TrimSpace(line2), 10, 32)

	result[0] = num
	result[1] = num2
	//line2, _ := reader.ReadString('\n')
	//splitted := strings.Split(line2, " ")

	//for i, elem := range splitted {
	//	x, _ := strconv.ParseInt(strings.TrimSpace(elem), 10, 64)
	//	result[i] = x
	//}

	var testNum int
	flag.IntVar(&testNum, "n", 0, "test number")
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)

	runner := y.Runner{}

	switch testNum {
	case 1:
		tc := &tt.TestRun{}
		result, err := runner.RunTestContest(tc, result...)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			writer.Flush()

			return

		}

		fmt.Fprintf(writer, "Result: ")

		for _, elem := range result {
			fmt.Fprintf(writer, "%d ", elem)
		}

	default:
		fmt.Fprintf(writer, "Incorrect test number")
	}

	writer.Flush()
}
