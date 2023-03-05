// https://contest.yandex.ru/contest/26365/problems/A/

package testrun

type TestRun struct {
}

func (*TestRun) Name() string {
	return "A+B"
}

func (*TestRun) RunInt(in ...int64) (result []int64, err error) {
	var res int64

	for idx := range in {
		res += in[idx]
	}

	result = append(result, res)

	return result, nil
}
