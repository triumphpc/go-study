package main

func alwaysFalse() bool {
	return false

}

func main() {

	switch alwaysFalse() {
	case true:
		println(true)
	case false:
		println(false)

	}

}
