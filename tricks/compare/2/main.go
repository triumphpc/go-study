package main

type data struct {
	num    int               // ok
	checks [10]func() bool   // несравниваемо
	doit   func() bool       // несравниваемо
	m      map[string]string // несравниваемо
	bytes  []byte            // несравниваемо
}

func main() {
	//v1 := data{}
	//v2 := data{}
	//fmt.Println("v1 == v2:",v1 == v2) //Нельзя сравнивать
}
