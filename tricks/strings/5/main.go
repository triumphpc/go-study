package main

import "log"

func main() {

	st := "c07a2b5d-7962-4a8c-87ce-bf445efe27cf.xml"

	ext := st[len(st)-3:]

	log.Print(ext)

	new := st[0 : len(st)-4]

	println(new)

}
