package main

func small() string {
	s := "hello, " + "world!"
	return s
}

func large() string {
	s := "a"
	s += "b"
	s += "c"
	s += "d"
	s += "e"
	s += "f"
	s += "g"
	s += "h"
	s += "i"
	s += "j"
	s += "k"
	s += "l"
	s += "m"
	s += "n"
	s += "o"
	s += "p"
	s += "q"
	s += "r"
	s += "s"
	s += "t"
	s += "u"
	s += "v"
	s += "w"
	s += "x"
	s += "y"
	s += "z"
	return s
}

func main() {
	small()
	large()
}
