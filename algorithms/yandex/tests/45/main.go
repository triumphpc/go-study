package main

import (
	"io"
	"os"
)

// Хэш таблица. Нативная реализация

func main() {
	task(os.Stdin, os.Stdout)
}

var pairs []int

func task(src io.Reader, dst io.Writer) {

}

func get(key int) int {
	for i := range pairs {
		if i == key {
			return pairs[key]
		}
	}

	return 0
}

func set(key, value int) bool {
	for i := range pairs {
		if i == key {
			pairs[i] = value

			return true
		}
	}

	return false
}
