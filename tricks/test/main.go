package main

type P struct {
	f string
}

func main() {
	var p *P
	if true {
		p = &P{}
	}

	print(p.f)
}
