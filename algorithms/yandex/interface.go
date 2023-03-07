package yandex

import mt "github.com/triumphpc/go-study/algorithms/metrics"

type TestContest interface {
	// RunInt - запуск скриптов с алгоритмом
	RunInt(in ...int) (result int, err error)
	// Name - наименования теста
	Name() string
}

type TestContestFloat interface {
	// Run - запуск скриптов с алгоритмом
	Run(k int, in ...float32) (result []float32, err error)
	// Name - наименования теста
	Name() string
}

type TestRunner interface {
	// RunTestContest - запускает тест
	RunTestContest(TestContest) (result int, err error)
}

type Runner struct {
}

func (*Runner) RunTestContest(tc TestContest, in ...int) (result int, err error) {
	defer mt.Duration(mt.Track(tc.Name()))

	return tc.RunInt(in...)
}

func (*Runner) RunTestContestFloat(tc TestContestFloat, k int, in ...float32) (result []float32, err error) {
	defer mt.Duration(mt.Track(tc.Name()))

	return tc.Run(k, in...)
}
