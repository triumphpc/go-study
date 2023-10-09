package main

import "fmt"

// Развернуть рекурсивно односвязанный список

type Item struct {
	pNext *Item
	val   rune
}

func createList() *Item {

	pHead := &Item{nil, 'a'}

	pCurr := pHead
	for i := 'b'; i <= 'z'; i++ {
		pItem := &Item{nil, i} // 1. Инициализация элемента со значением
		pCurr.pNext = pItem    // 2. Назначаем этот элемент как следующий для текущего
		pCurr = pItem          // 3. Присваиваем текущему данный элемент
	}

	return pHead // 4. Возвращаем голову
}

func printList(pList *Item) {

	pCurr := pList
	for {
		fmt.Printf("%c", pCurr.val)

		if pCurr.pNext != nil {
			pCurr = pCurr.pNext
		} else {
			break
		}
	}
	fmt.Println("")
}

func reverseList(pList *Item) *Item {

	pCurr := pList       // 1. Текущий элемент = голове списка
	var pTop *Item = nil // 2. Определяем последний
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.pNext // 3. Сохраняем следующий от текущего
		pCurr.pNext = pTop   // 4. Текущему проставляем следующий (от предыдущего шага)
		pTop = pCurr         // 5. Запомним текущий
		pCurr = pTemp        // 6. Текущий = следующий от текущего
	}

	return pTop
}

func main() {

	var pList = createList()
	printList(pList)
	fmt.Println()
	printList(reverseList(pList))
}
