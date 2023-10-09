package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Два велосипеда
// Вася решил накопить денег на два одинаковых велосипеда — себе и сестре. У Васи есть копилка, в которую каждый день
// он может добавлять деньги (если, конечно, у него есть такая финансовая возможность). В процессе накопления Вася не вынимает деньги из копилки.
//У вас есть информация о росте Васиных накоплений — сколько у Васи в копилке было денег в каждый из дней.
//Ваша задача — по заданной стоимости велосипеда определить
//* первый день, в которой Вася смог бы купить один велосипед,
//* и первый день, в который Вася смог бы купить два велосипеда.
//Подсказка: решение должно работать за O(log n).
// В первой строке дано число дней n, по которым велись наблюдения за Васиными накоплениями. 1 ≤ n ≤ 106.
//В следующей строке записаны n целых неотрицательных чисел. Числа идут в порядке неубывания. Каждое из чисел не превосходит 106.
//В третьей строке записано целое положительное число s — стоимость велосипеда. Это число не превосходит 106.

func main() {
	task2(os.Stdin, os.Stdout)

}

var step1, step2 = -1, -1
var cost, cost2 int
var line []string

func task(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		_ = writer.Flush()
	}()
	scanner := bufio.NewScanner(src)

	// num of days
	scanner.Scan()
	num := conv(scanner.Text())

	// rates
	scanner.Scan()
	line = strings.Split(scanner.Text(), " ")

	scanner.Scan()
	cost = conv(scanner.Text())
	cost2 = cost * 2

	defer func() {
		_, _ = writer.WriteString(strconv.Itoa(step1))
		_, _ = writer.WriteString(strconv.Itoa(step2))
	}()

	rec(0, num, &cost, &step1)
	rec(cost, num, &cost2, &step2)
}

func conv(n string) int {
	v, _ := strconv.Atoi(n)

	return v
}

func rec(left, right int, cost, step *int) {
	if left == right { // 1. Сравниваем левую и правую границу
		return
	}

	mid := (left + right) / 2 // 2. Берем середину
	if mid == left {
		return
	}

	cur := conv(line[mid])
	if cur >= *cost { // 3. Если больше и равно - прибавляем шаг
		*step = mid + 1
		rec(left, mid, cost, step)
	} else { // 4. В противном случае сужаем поиск на нужную срез
		rec(mid, right, cost, step)
	}
}

func findDayByPrice(from int, to int, values []int, price int) int {
	mid := from + (to-from)/2
	// осталось 2 элемента
	if to-from == 1 {
		if values[from] >= price {
			return from + 1
		}
		if to == len(values) || values[to] < price {
			return -1
		}
		return to + 1
	}
	if values[mid] < price {
		return findDayByPrice(mid, to, values, price)
	}
	return findDayByPrice(from, mid, values, price)
}
func task2(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	days, _ := strconv.Atoi(scanner.Text())
	values := make([]int, 0, days)
	for i := 0; i < days; i++ {
		scanner.Scan()
		intVal, _ := strconv.Atoi(scanner.Text())
		values = append(values, intVal)
	}
	scanner.Scan()
	price, _ := strconv.Atoi(scanner.Text())
	from := 0
	to := days
	double := findDayByPrice(from, to, values, price*2)
	if double != -1 {
		to = double
	}
	fmt.Fprintf(dst, "%d %d\n", findDayByPrice(from, to, values, price), double)
}
