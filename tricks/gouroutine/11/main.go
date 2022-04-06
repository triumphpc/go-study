package main

import "sync"

// Когда вы определяете новый тип на основе существующего (не интерфейсного), тем самым вы создаёте объявление
//типа и не наследуете методы, объявленные в существующем типе.

//type myMutex sync.Mutex

//func main() {
//	var mtx myMutex
//	mtx.Lock() // ошибка
//	mtx.Unlock() // ошибка
//}

// Если вам нужны методы из исходного типа, вы можете задать новый тип структуры, встроив исходный в качестве анонимного поля.
type myLocker struct {
	sync.Mutex
}

func main() {
	var lock myLocker
	lock.Lock()   // ok
	lock.Unlock() // ok
}

// Объявления интерфейсных типов также сохраняют свои наборы методов.

//type myLocker sync.Locker
//
//func main() {
//	var lock myLocker = new(sync.Mutex)
//	lock.Lock() // ok
//	lock.Unlock() // ok
//}
