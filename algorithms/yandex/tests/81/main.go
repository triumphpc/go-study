// Вы хотите набрать строку 𝑠
//, состоящую из 𝑛
// строчных латинских букв, в вашем любимом текстовом редакторе Notepad#.
//
//Notepad# поддерживает два типа операций:
//
//дописать любую букву в конец строки;
//скопировать непрерывную подстроку уже напечатанной строки и вставить эту подстроку в конец строки.
//Можете ли вы набрать строку 𝑠
// строго меньше, чем за 𝑛
// операций?
//
//Входные данные
//В первой строке записано одно целое число 𝑡
// (1≤𝑡≤104
//) — количество наборов входных данных.
//
//В первой строке каждого набора входных данных записано одно целое число 𝑛
// (1≤𝑛≤2⋅105
//) — длина строки 𝑠
//.
//
//Во второй строке записана строка 𝑠
//, состоящая из 𝑛
// строчных латинских букв.
//
//Сумма 𝑛
// по всем наборам входных данных не превосходит 2⋅105
//.
//
//Выходные данные
//На каждый набор входных данных выведите «YES», если вы можете набрать строку 𝑠
// строго меньше, чем за 𝑛
// операций. В противном случае выведите «NO».

// Пример
//входные данныеСкопировать
//6
//10
//codeforces
//8
//labacaba
//5
//uohhh
//16
//isthissuffixtree
//1
//x
//4
//momo
//выходные данныеСкопировать
//NO
//YES
//NO
//YES
//NO
//YES
//Примечание
//В первом наборе входных данных можно начать с набора «codef» (5
// операций), затем скопировать «o» (1
// операций) из уже набранной части, затем завершить набором «rces» (4
// операции). Это будет 10
// операций, что не строго меньше 𝑛
//. Существуют и другие способы набрать «codeforces». Однако, что бы вы ни делали, нельзя это сделать меньше, чем за 𝑛
// операций.
//
//Во втором наборе входных данных можно набрать «labac» (5
// операций), затем скопировать «aba» (1
// операция), завершив строку за 6
// операций.
// https://codeforces.com/problemset/problem/1766/B

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	// 1. Сохраняем в хэщ строки большие 1 - каждую подстроку
	// 2. Если в мапе есть текущая+1, проверяем езе плюс 1 и тд и подсчитываем количество операций

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	//res := make(map[string][]int)
	scanner.Scan()
	world := scanner.Text()

	chrs := make(map[string]bool, len(world))
	chrs[world[0:2]] = true // Первую подстроку сраззу положим в кэщ
	cnt := 2 + len(world)%2 // Первые два символа это две операции уже + остаток от целочисленого деления на два

	for left := 2; left < len(world); left++ { // Далее по каждому проеряем была ли в кэше уже поддстрока или нет
		// Внутренний цикл проверки
		right := left + 2 // Правый маркер
		if right > len(world)-1 {
			break
		}

		for {
			if chrs[world[left:right]] { // Если уже есть, значит двигаем дальше
				right++ // Увеличиываем шаг, пока не встретим вхождения
			} else {
				chrs[world[0:right]] = true // Запомним текущую подстроку и с самого начала
				chrs[world[left:right]] = true
				cnt++ // Увеличиваем количество операций

				break
			}
		}
	}

	fmt.Println(cnt, chrs)
}