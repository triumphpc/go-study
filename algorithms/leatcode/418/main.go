// https://leetcode.com/problems/sentence-screen-fitting/

package main

import "log"

// Если задан экран rows x cols и предложение, представленное в виде списка строк, верните количество раз,
//которое данное предложение может быть помещено на экран. Порядок слов в предложении должен оставаться неизменным,
//и слово не может быть разбито на две строки. Два последовательных слова в строке должны разделяться одним пробелом.

func wordsTyping(sentence []string, rows int, cols int) int {
	sentenceStr := ""
	for _, word := range sentence {
		sentenceStr += word + " "
	}
	length := len(sentenceStr)
	count := 0

	for i := 0; i < rows; i++ {
		count += cols

		if sentenceStr[count%length] == ' ' {
			count++
		} else {
			for count > 0 && sentenceStr[(count-1)%length] != ' ' {
				count--
			}
		}
	}

	return count / length
}

func main() {
	sentence := []string{"12345678", ""}
	rows := 2
	cols := 8

	result := wordsTyping(sentence, rows, cols)

	log.Println(result)

}
