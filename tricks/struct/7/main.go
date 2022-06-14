package main

import "fmt"

// Какие типы не являются встраиваемыми
type (
	A *int // ссылочный тип не является встраиваемым
	B = *int
	C = interface{}
	D = C
	E = *C // нет
	F = *B // нет
	G = struct{}
	H = *G
	K *G
)

type T struct {
	//A
	B
	C
	D
	//E
	//F
	G
	H
	//K
}

func main() {
	fmt.Println(T{})
	fmt.Printf("%+T\n", new(A)) // *main.A
	fmt.Printf("%+T\n", new(B)) // **int
	fmt.Printf("%+T\n", new(C)) // *interface {}
	fmt.Printf("%+T\n", new(D)) // *interface {}
	fmt.Printf("%+T\n", new(E)) // **interface {}
	fmt.Printf("%+T\n", new(F)) // ***int
	fmt.Printf("%+T\n", new(G)) // *struct{}
	fmt.Printf("%+T\n", new(H)) // **struct{}
	fmt.Printf("%+T\n", new(K)) // *main.K

}
