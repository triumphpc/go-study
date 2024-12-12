package main

import (
	"fmt"
)

// Задача:
// Есть сервис обработки заказов. Он должен принимать заказы,
// обрабатывать их и возвращать результат.
// Найдите ошибку в реализации.

type Order struct {
	ID   int
	Data string
}

func processOrders() {
	orders := make(chan Order)   // канал для заказов
	results := make(chan string) // канал для результатов

	// Горутина для обработки заказов
	go func() {
		for order := range orders {
			// Обработка заказа
			result := fmt.Sprintf("Заказ %d обработан", order.ID)
			results <- result
		}
	}()

	// Основной цикл
	for i := 1; i <= 3; i++ {
		orders <- Order{ID: i, Data: "данные"}
		//fmt.Println(<-results)
	}

	close(orders)

	close(results)
}

func main() {
	processOrders()
}
