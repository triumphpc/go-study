package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Скобочная последовательность
// Вот какую задачу Тимофей предложил на собеседовании одному из кандидатов. Если вы с ней ещё не сталкивались, то наверняка столкнётесь –— она довольно популярная.
//Дана скобочная последовательность. Нужно определить, правильная ли она.
//Будем придерживаться такого определения:
//* пустая строка —– правильная скобочная последовательность;
//* правильная скобочная последовательность, взятая в скобки одного типа, –— правильная скобочная последовательность;
//* правильная скобочная последовательность с приписанной слева или справа правильной скобочной последовательностью —– тоже правильная.
//На вход подаётся последовательность из скобок трёх видов: [], (), {}.
//Напишите функцию is_correct_bracket_seq, которая принимает на вход скобочную последовательность и возвращает True, если последовательность правильная, а иначе False.

func main() {
	task(os.Stdin)

}

const (
	strScope1 = 40 // 1. Определяем символы скобок
	endScope1 = 41
	strScope2 = 123
	endScope2 = 125
	strScope3 = 91
	endScope3 = 93
)

func task(src io.Reader) {
	fmt.Println(checkSequence(src))
}

func checkSequence(src io.Reader) bool {
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	slice := []byte(scanner.Text())
	stack := make([]int8, 0, len(slice)/2)
	mapIdx := map[int8]int8{endScope1: strScope1, endScope2: strScope2, endScope3: strScope3} // Определяем мапу соответсвий символов

	for idx := range slice {
		switch slice[idx] {
		case endScope2, endScope1, endScope3: // Если закрытый символ, то смотрим а скольок было входных символов
			if len(stack) == 0 && stack[len(stack)-1] != mapIdx[int8(slice[idx])] {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, int8(slice[idx]))
		}
	}

	return len(stack) == 0
}
