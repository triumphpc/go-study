package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Эффективная быстрая сортировка
// Тимофей решил организовать соревнование по спортивному программированию, чтобы найти талантливых стажёров. Задачи подобраны,
//участники зарегистрированы, тесты написаны. Осталось придумать, как в конце соревнования будет определяться победитель.
//Каждый участник имеет уникальный логин. Когда соревнование закончится, к нему будут привязаны два показателя:
//количество решённых задач Pi и размер штрафа Fi. Штраф начисляется за неудачные попытки и время, затраченное на задачу.
//
//Тимофей решил сортировать таблицу результатов следующим образом: при сравнении двух участников выше будет идти тот,
//у которого решено больше задач. При равенстве числа решённых задач первым идёт участник с меньшим штрафом.
//Если же и штрафы совпадают, то первым будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке.
//
//Тимофей заказал толстовки для победителей и накануне поехал за ними в магазин. В своё отсутствие он поручил вам
//реализовать алгоритм быстрой сортировки (англ. quick sort) для таблицы результатов. Так как Тимофей любит спортивное
//программирование и не любит зря расходовать оперативную память, то ваша реализация сортировки не может
//потреблять O(n) дополнительной памяти для промежуточных данных (такая модификация быстрой сортировки называется "in-place").
//
//
//
//Как работает in-place quick sort
//
//Как и в случае обычной быстрой сортировки, которая использует дополнительную память, необходимо выбрать
//опорный элемент (англ. pivot), а затем переупорядочить массив. Сделаем так, чтобы сначала шли элементы,
//не превосходящие опорного, а затем —– большие опорного.
//Затем сортировка вызывается рекурсивно для двух полученных частей. Именно на этапе разделения элементов
//на группы в обычном алгоритме используется дополнительная память. Теперь разберёмся, как реализовать
//этот шаг in-place.
//Пусть мы как-то выбрали опорный элемент. Заведём два указателя left и right, которые изначально
//будут указывать на левый и правый концы отрезка соответственно. Затем будем двигать левый
//указатель вправо до тех пор, пока он указывает на элемент, меньший опорного. Аналогично двигаем
//правый указатель влево, пока он стоит на элементе, превосходящем опорный. В итоге окажется, что что левее
//от left все элементы точно принадлежат первой группе, а правее от right — второй. Элементы, на которых
//стоят указатели, нарушают порядок. Поменяем их местами (в большинстве языков программирования используется
//функция swap()) и продвинем указатели на следующие элементы. Будем повторять это действие до тех пор, пока
//left и right не столкнутся.

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text()) // 1. Размер

	competitors := make([]*competitor, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		login := scanner.Text() // 2. логин 10 5
		scanner.Scan()
		completedTasks, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		fine, _ := strconv.Atoi(scanner.Text())

		competitors = append(competitors, &competitor{
			login:         login,
			completeCount: completedTasks,
			fineCount:     fine,
		})
	}

	quickSort(competitors, 0, size-1)

	fmt.Println("--------")

	for idx := range competitors {
		fmt.Println(competitors[idx].login, competitors[idx].completeCount, competitors[idx].fineCount)
	}
}

type competitor struct {
	login         string
	completeCount int
	fineCount     int
}

func (c *competitor) lessCompetitor(b *competitor) bool { // 2 Реализация функции сравнения
	if c.completeCount > b.completeCount {
		return true
	}

	if c.completeCount == b.completeCount {
		if c.fineCount > b.fineCount {
			return true
		}

		if c.fineCount == b.fineCount {
			return c.login > b.login
		}
	}

	return false

}

func quickSort(arr []*competitor, left, right int) {
	if right-left <= 1 {
		if arr[right].lessCompetitor(arr[left]) {
			arr[left], arr[right] = arr[right], arr[left]
		}

		return
	}

	mid := left + (right-left)/2 // 3. Опорный элемент
	pivot := arr[mid]
	leftPt := left // 4. Левый и правый указатели
	rightPt := right

	for leftPt < rightPt {
		if arr[leftPt].lessCompetitor(pivot) {
			leftPt++ // Двигаем указатели
			continue
		}

		if pivot.lessCompetitor(arr[rightPt]) {
			rightPt--
			continue
		}

		// swap
		arr[leftPt], arr[rightPt] = arr[rightPt], arr[leftPt] // 5. Меняем местами
		leftPt++
		rightPt--
	}
	// 6. Определяем новое положение опорного элемента
	if pivot == arr[leftPt] {
		mid = leftPt
	} else {
		mid = rightPt
	}

	quickSort(arr, left, mid)
	quickSort(arr, mid, right)
}
