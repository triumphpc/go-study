package yandex

import mt "github.com/triumphpc/go-study/algorithms/metrics"

type TestContest interface {
	// RunInt - запуск скриптов с алгоритмом
	RunInt(in ...int64) (result []int64, err error)
	// Name - наименования теста
	Name() string
}

type TestRunner interface {
	// RunTestContest - запускает тест
	RunTestContest(TestContest) (result []int64, err error)
}

type Runner struct {
}

func (*Runner) RunTestContest(tc TestContest, in ...int64) (result []int64, err error) {
	defer mt.Duration(mt.Track(tc.Name()))
	return tc.RunInt(in...)
}
