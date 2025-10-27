package _25_valid_palindrome

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := "0P"

	k := isPalindrome(input)
	fmt.Println(k)

}

func isPalindrome(s string) bool {
	//s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left <= right {
		// Пропускаем не-буквы слева
		if !isLetter(s[left]) {
			left++
			continue
		}

		// Пропускаем не-буквы справа
		if !isLetter(s[right]) {
			right--
			continue
		}

		// Сравниваем буквы
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z'
}
