package main

import (
	"flag"
	"fmt"
	"os"

	y "github.com/triumphpc/go-study/algorithms/yandex"
	tt "github.com/triumphpc/go-study/algorithms/yandex/tests/1"
	tt2 "github.com/triumphpc/go-study/algorithms/yandex/tests/2"
	tt3 "github.com/triumphpc/go-study/algorithms/yandex/tests/3"
	tt4 "github.com/triumphpc/go-study/algorithms/yandex/tests/4"
	tt5 "github.com/triumphpc/go-study/algorithms/yandex/tests/5"
	tt6 "github.com/triumphpc/go-study/algorithms/yandex/tests/6"
)
import "bufio"
import "strings"
import "strconv"

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := make([]int, 2)

	line1, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(line1), 10, 32)
	splitted := strings.Split(line1, " ")

	resultSplitted := make([]float32, len(splitted))
	for i, elem := range splitted {
		x, _ := strconv.ParseFloat(strings.TrimSpace(elem), 10)
		resultSplitted[i] = float32(x)
	}

	line2, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseInt(strings.TrimSpace(line2), 10, 32)

	splitted = strings.Split(line2, " ")
	resultSplitted2 := make([]float32, len(splitted))
	for i, elem := range splitted {
		x, _ := strconv.ParseFloat(strings.TrimSpace(elem), 10)
		resultSplitted2[i] = float32(x)
	}

	result[0] = int(num)
	result[1] = int(num2)

	var testNum int
	flag.IntVar(&testNum, "n", 0, "test number")
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)
	runner := y.Runner{}

	var err error
	var data interface{}

	switch testNum {
	case 1:
		tc := &tt.TestRun{}
		data, err = runner.RunTestContest(tc, result...)
	case 2:
		tc := &tt2.TestRun{}
		data, err = runner.RunTestContestFloat(tc, int(num2), resultSplitted...)
	case 3:
		tc := &tt3.TestRun{}
		data, err = runner.RunTestContestFloat(tc, int(num2), resultSplitted...)
	case 4:
		tc := &tt4.TestRun{}
		data, err = runner.RunTestContestFloat(tc, int(num2), resultSplitted...)
	case 5:
		tc := &tt5.TestRun{}
		data, err = runner.RunTestContestFloat(tc, int(num2), resultSplitted...)
	case 6:
		tc := &tt6.TestRun{}
		data, err = runner.RunTestContestFloat(tc, int(num2), resultSplitted...)

	default:
		fmt.Fprintf(writer, "Incorrect test number")

		return
	}

	if err != nil {
		fmt.Fprintf(writer, err.Error())
		writer.Flush()

		return
	}

	fmt.Fprintf(writer, "Result: ")
	fmt.Fprintf(writer, "%v ", data)

	writer.Flush()
}
