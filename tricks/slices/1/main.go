package main

import (
	"os"
	"regexp"
)

func main() {

}

var digitsRegexp := regexp.MustCompile("[0-9]+")

func FindDigits(fileName string) []byte {
	b, _ := os.ReadFile(fileName)

	 return digitsRegexp.Find(b) // Проблема: создаст слайс (не большой) с ссылкой на базовый массив огромной длинны
	 // Решается создаем нового слайса и copy
	 // res := make([]byte, len(b))
	 // copy(res, b)
	 // return res
}
