package main

import "fmt"

// Какие типы не являются встраиваемыми?

type (
	A *int          // ссылочный тип не является встраиваемым
	B = *int        // указатель на int можно
	C = interface{} // интерфейс можно
	D = C           // // это клон - можно
	E = *C          // указатель на интерфейс нельзя
	F = *B          // Значение B = указатель на int, нельзя
	G = struct{}    // Можно
	H = *G          // H =  указатель на G структуру- можно
	K *G            // ссылочный тип нельзя
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
