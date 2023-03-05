package main

type T struct {
	_ int // Нельзя присваивать
	_ bool
}

func main() {
	//var t1 = T{123, true} ошибка, нельзя присваивать пустым полям
	//var t2 = T{126, false}

	//fmt.Println(t1 == t2)

}
