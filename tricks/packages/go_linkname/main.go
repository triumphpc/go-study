// Директива go:linkname - это специальная директива компилятора Go, которая позволяет связать символ (функцию или переменную)
// из одного пакета с символом из другого пакета, даже если он не экспортирован (начинается с маленькой буквы).

// Важные моменты:
//
//Директива должна использоваться вместе с импортом "unsafe"
//Между //go:linkname и объявлением функции не должно быть пустых строк
//Требует особой осторожности, так как обходит систему типов Go
//
//Вот несколько примеров:

package runtime

import (
	"unsafe"
	_ "unsafe"
)

//go:linkname gopark runtime.gopark
func gopark(unlockf func(*g, unsafe.Pointer) bool, lock unsafe.Pointer, reason waitReason, traceEv byte, traceskip int)

//go:linkname goready runtime.goready
func goready(gp *g, traceskip int)

func CustomPark() {
	gopark(nil, nil, waitReasonCustom, traceEvGoBlock, 1)
}
