package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Подпоследовательность
// Дана строка s, найдите длину самой длинной подстрока без повторения символов.
// dvdf
// dvdfdfkloi
// https://www.code-recipe.com/post/longest-substring-without-repeating-characters

// O(n) - сложность алгоритма O(1) - сложность по памяти, так как работа с мапой
// Мы проверяем каждый символ в строке только один раз. Да, нам также придется удалить элементы из хэш-карты, как
// только будет найден повторяющийся символ, но это операция с постоянным временем, потому что в нашем случае карта может
// хранить только 62 элемента до того, как будет найден дубликат.

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) { // On
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	s := scanner.Text()

	fmt.Fprintf(dst, "%d ", lengthOfLongestSubstring1(s))
}

// Наивный алгоритм 1
func lengthOfLongestSubstring1(s string) int {
	sBites := []byte(s)
	var max int

	for i := 0; i < len(sBites); i++ {
		hs := make(map[byte]bool, len(sBites)-i)
		cur := 0

		for j := i; j < len(sBites); j++ {
			if hs[sBites[j]] {
				break
			} else {
				hs[sBites[j]] = true
				cur++
			}
		}

		if cur > max {
			max = cur
		}
	}

	return max
}

// Наивный алгоритм 2
func lengthOfLongestSubstring2(s string) int {
	n := len(s)

	// String with one character is always unique, return 1
	if n == 1 {
		return 1
	}

	var result int
	for start := 0; start < n-1; start++ {
		for end := start; end < n; end++ {
			// Hashmap to store characters of the current substring
			// Needed to check if all the characters are unique or if it contains duplicates
			characterMap := make(map[uint8]bool)

			// True if substring contains all unique characters
			// False if it contains duplicates
			isUniqueSubstring := true

			// Loop to check if the current substring contains unique characters or has duplicates
			for i := start; i <= end; i++ {
				if _, isPresent := characterMap[s[i]]; !isPresent {
					characterMap[s[i]] = true
					continue // continue to next character in the substring
				}

				// substring contains duplicates
				isUniqueSubstring = false
				break
			}

			// If the current substring has all unique characters
			// Update result if current substring length is greater than result so far
			if isUniqueSubstring {
				result = max(result, end-start+1)
			}
		}
	}
	return result
}

// max returns the max of num1 and num2
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// Метод скользящего окна. Самый быстрый
func lengthOfLongestSubstring3(s string) int {
	// Определение длинны строки
	n := len(s)

	// Результат количества
	var result int

	// Map to store visited characters along with their index
	// Хэш истории индекса для символов
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	// Стартовый индекс текущей подстроки
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	// Начинаем проходить строку пошагово
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		// Проверяем, был ли символ уже ранее
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			// Максимальное из текущего значения или текущий шаг-стартовый для подстроки
			result = max(result, end-start)

			// Remove all characters before the new
			// Удаляем все подстроки из истории до настроящего шага
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			// Двигаем старт подстроки до дубликата + 1
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		// Добавляем символ текущей итерации в историю
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	// Для случая, если вообще не было дубликатов - полный размер строки
	result = max(result, n-start)

	return result

}
