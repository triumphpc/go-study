// https://contest.yandex.ru/contest/26365/problems/A/
// 1536 Kb vs 384
// 1024 Kb vs 384
// ------
// 640 minus

package testrun

type TestRun struct {
}

func (*TestRun) Name() string {
	return "A+B"
}

func (*TestRun) RunInt(in ...int) (result int, err error) {
	for idx := range in {
		result += in[idx]
	}
	//
	//result = append(result, res)
	//
	//return result, nil

	return result, nil
}
