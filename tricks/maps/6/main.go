package main

import "fmt"

type T struct {
	cn     string
	street string
}

type View struct {
	ExpectedDocs map[string]map[string]*ViewDocMeta
}

type ViewDocMeta struct {
	EventID string //  ИД события с досылом документа
}

type RequestMetaPG struct {
	*View
}

func main() {
	//names := []string{"kasi", "remya", "nandan"}
	//
	//m := make(map[string]map[string]T, len(names))
	//for _, name := range names {
	//
	//	//t := T{cn: "Chaithra", street: "fkmp"}
	//
	//	if m["uid"] == nil {
	//		m["uid"] = map[string]T{}
	//	}
	//
	//	m["uid"][name] = T{cn: "Chaithra", street: "fkmp"}
	//
	//}
	//fmt.Println(m)
	//
	//y, _ := yaml.Marshal(&m)
	//
	//fmt.Println(string(y))
	////fmt.Println(m, names)
	//
	//
	//ExpectedDocs  map[string]map[string]*ViewDocMeta

	//v := &View{}
	//r := &RequestMetaPG{v}
	//
	//if r.ExpectedDocs == nil {
	//	r.ExpectedDocs = map[string]map[string]*ViewDocMeta{}
	//}
	//
	//if r.ExpectedDocs["fff"] == nil {
	//	r.ExpectedDocs["fff"] = make(map[string]*ViewDocMeta, 2)
	//	fmt.Println("ffff")
	//
	//}

	list := [...]string{}
	fmt.Println(cap(list), len(list))

}
