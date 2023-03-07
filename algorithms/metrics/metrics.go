package metrics

// go tool pprof -http=":9090" /tmp/mem.profile
// go tool pprof -http=":9090" /tmp/cpu.profile

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var fmem, fcpu *os.File

func Track(msg string) (string, time.Time) {
	// Журнал профилирование CPU
	var err error
	//fcpu, err = os.Create(`/tmp/cpu.profile`)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := pprof.StartCPUProfile(fcpu); err != nil {
	//	panic(err)
	//}

	// создаём файл журнала профилирования памяти
	fmem, err = os.Create("/tmp/mem.profile")
	if err != nil {
		panic(err)
	}

	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	//defer fcpu.Close()
	defer fmem.Close()
	defer pprof.StopCPUProfile()

	log.Printf("%v: %v\n", msg, time.Since(start))

	// Снимаем слепок по памяти
	runtime.GC() // получаем статистику по использованию памяти
	if err := pprof.WriteHeapProfile(fmem); err != nil {
		panic(err)
	}
}
