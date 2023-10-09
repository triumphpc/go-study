package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Хэш таблица
// Тимофей, как хороший руководитель, хранит информацию о зарплатах своих сотрудников в базе данных и постоянно её обновляет. Он поручил вам написать реализацию хеш-таблицы, чтобы хранить в ней базу данных с зарплатами сотрудников.
// Хеш-таблица должна поддерживать следующие операции:
// * put key value —– добавление пары ключ-значение. Если заданный ключ уже есть в таблице, то соответствующее ему значение обновляется.
// * get key –— получение значения по ключу. Если ключа нет в таблице, то вывести «None». Иначе вывести найденное значение.
// * delete key –— удаление ключа из таблицы. Если такого ключа нет, то вывести «None», иначе вывести хранимое по данному ключу значение и удалить ключ.
// В таблице хранятся уникальные ключи.
// Требования к реализации:
// * Нельзя использовать имеющиеся в языках программирования реализации хеш-таблиц (std::unordered_map в С++, dict в Python, HashMap в Java, и т. д.)
// * Число хранимых в таблице ключей не превосходит 105.
// * Разрешать коллизии следует с помощью метода цепочек или с помощью открытой адресации.
// * Все операции должны выполняться за O(1) в среднем.
// * Поддерживать рехеширование и масштабирование хеш-таблицы не требуется.
// * Ключи и значения, id сотрудников и их зарплата, —– целые числа. Поддерживать произвольные хешируемые типы не требуется.
// Формат ввода
//
// В первой строке задано общее число запросов к таблице n (1≤ n≤ 106).
// В следующих n строках записаны запросы, которые бывают трех видов –— get, put, delete —– как описано в условии.
// Все ключи и значения –— целые неотрицательные числа, не превосходящие 109.
// Формат вывода
//
// На каждый запрос вида get и delete выведите ответ на него в отдельной строке.
func main() {
	task(os.Stdin, os.Stdout)
}

type bucketChain struct {
	key  int
	val  int
	next *bucketChain
}
type hashTableInt struct {
	buckets []*bucketChain
	m       int // количество
}

func (ht *hashTableInt) bucketID(key int) int {
	return key % ht.m
}

func newHashTableInt(m int) *hashTableInt {
	return &hashTableInt{
		m:       m,
		buckets: make([]*bucketChain, m),
	}
}

func (ht *hashTableInt) get(key int) (value int, isExists bool) {
	bucketNum := ht.bucketID(key)
	bChain := ht.buckets[bucketNum]
	// ищем в цепи ключ и возвращаем значение
	for bChain != nil {
		if bChain.key == key {
			return bChain.val, true
		}
		bChain = bChain.next
	}
	return
}

func (ht *hashTableInt) put(key, val int) {
	bucketNum := ht.bucketID(key)
	bChain := ht.buckets[bucketNum]

	if bChain == nil {
		ht.buckets[bucketNum] = &bucketChain{key, val, nil}

		return
	}

	// ищем в цепи ключ и обновляем
	for bChain != nil {
		if bChain.key == key {
			bChain.val = val

			return

		}

		bChain = bChain.next

	}

	// если не нашли ключ, создаем элемент цепи в HEAD

	ht.buckets[bucketNum] = &bucketChain{key, val, ht.buckets[bucketNum]}

}

func (ht *hashTableInt) delete(key int) (val int, isExist bool) {
	bucketNum := ht.bucketID(key)
	bChain := ht.buckets[bucketNum]

	var prevElem *bucketChain

	for bChain != nil {
		if bChain.key == key {
			// исключаем из цепи, если есть предыдущий
			if prevElem != nil {
				prevElem.next = bChain.next
				// иначе просто выставляем nil для корзины
			} else {
				ht.buckets[bucketNum] = nil
			}

			return bChain.val, true
		}

		prevElem = bChain
		bChain = bChain.next
	}

	return
}

const M = 2 // Количество бакетов

func task(src io.Reader, dst io.Writer) {
	table := newHashTableInt(M)

	scanner := bufio.NewScanner(src)
	scanner.Scan()

	commandsTotal, _ := strconv.Atoi(scanner.Text())
	for commandsTotal > 0 {
		scanner.Scan()
		commandParts := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		switch commandParts[0] {
		case "get":
			key, _ := strconv.Atoi(commandParts[1])
			value, isExists := table.get(key)
			if !isExists {
				fmt.Fprint(dst, "None\n")
			} else {
				fmt.Fprintf(dst, "%d\n", value)
			}
		case "put":
			key, _ := strconv.Atoi(commandParts[1])
			value, _ := strconv.Atoi(commandParts[2])
			table.put(key, value)
		case "delete":
			key, _ := strconv.Atoi(commandParts[1])
			value, isExists := table.delete(key)
			if !isExists {
				fmt.Fprint(dst, "None\n")
			} else {
				fmt.Fprintf(dst, "%d\n", value)
			}
		}
		commandsTotal--
	}

}
