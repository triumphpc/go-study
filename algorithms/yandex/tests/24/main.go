package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Генерация скобок
//Рита по поручению Тимофея наводит порядок в правильных скобочных последовательностях (ПСП), состоящих только из круглых скобок (). Для этого ей надо сгенерировать все ПСП длины 2n в алфавитном порядке —– алфавит состоит из ( и ) и открывающая скобка идёт раньше закрывающей.
//Помогите Рите —– напишите программу, которая по заданному n выведет все ПСП в нужном порядке.
//Рассмотрим второй пример. Надо вывести ПСП из четырёх символов. Таких всего две:
//1. (())
//2. ()()
//(()) идёт раньше ()(), так как первый символ у них одинаковый, а на второй позиции у первой ПСП стоит (, который идёт раньше ).
//
// На вход принимаем число - количество скобок

func main() {
	task(os.Stdin, os.Stdout)

}

func task(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	scanner.Scan()
	num := conv(scanner.Text())

	genBrackets("", num, 0, 0, dst)

}

func conv(n string) int {
	v, _ := strconv.Atoi(n)

	return v
}

func genBrackets(result string, num, left, right int, dst io.Writer) {
	if left == num && right == num { // 1. Если левый и правый границы равны - выводим
		fmt.Println("0")
		fmt.Fprintf(dst, "%s\n", result)
		return
	}
	if left != num { // 2. Если левый шаг не дошел до нужного количества - смещаем его и в рекурсию
		fmt.Println("X", left, num)
		genBrackets(result+"(", num, left+1, right, dst)
		fmt.Println("OUT X", left, num)
	}
	if left > right { // 3. Далее отработка правой части, пока не заполнит все остатки
		fmt.Println("Y", left, right)
		genBrackets(result+")", num, left, right+1, dst)

		fmt.Println("OUT Y", left, num)
	}
}
