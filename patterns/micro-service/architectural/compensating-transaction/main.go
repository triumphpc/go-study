// Паттерн "Compensating Transaction" используется для управления сложными транзакциями в распределенных системах,
//где невозможно использовать традиционные механизмы транзакций, такие как ACID. Этот паттерн помогает откатить изменения,
//если одна из операций в цепочке транзакций не удалась, обеспечивая целостность данных.

// Проблемы, которые решает Compensating Transaction
//Отказоустойчивость: В распределенных системах, где транзакции могут затрагивать несколько сервисов, важно иметь возможность откатить изменения, если что-то пошло не так.
//Целостность данных: Обеспечивает согласованность данных, даже если одна из операций в транзакции не удалась.
//Гибкость: Позволяет управлять транзакциями в системах, где невозможно использовать традиционные транзакционные механизмы.

//Плюсы Compensating Transaction
//Целостность данных: Обеспечивает согласованность данных в случае сбоя.
//Гибкость: Подходит для сложных распределенных систем, где невозможно использовать ACID-транзакции.
//Масштабируемость: Позволяет управлять транзакциями в масштабируемых микросервисных архитектурах.

// Минусы Compensating Transaction
// Сложность реализации: Требует тщательного проектирования и реализации логики компенсации.
// Задержки: Может увеличить время выполнения транзакции из-за необходимости выполнения компенсационных операций.
// Неопределенность: В некоторых случаях может быть сложно определить, как именно откатить изменения.
package main

import (
	"errors"
	"fmt"
)

// BookingStep - интерфейс для шагов бронирования
type BookingStep interface {
	Execute() error
	Compensate() error
}

// ReserveSeat - шаг резервирования места
type ReserveSeat struct {
	seatID string
}

func (r *ReserveSeat) Execute() error {
	fmt.Println("Reserving seat:", r.seatID)
	// Симуляция успешного выполнения
	return nil
}

func (r *ReserveSeat) Compensate() error {
	fmt.Println("Releasing seat:", r.seatID)
	// Симуляция успешного отката
	return nil
}

// ProcessPayment - шаг обработки платежа
type ProcessPayment struct {
	amount float64
}

func (p *ProcessPayment) Execute() error {
	fmt.Println("Processing payment of:", p.amount)
	// Симуляция ошибки
	return errors.New("payment failed")
}

func (p *ProcessPayment) Compensate() error {
	fmt.Println("Refunding payment of:", p.amount)
	// Симуляция успешного отката
	return nil
}

// ExecuteTransaction - выполняет транзакцию с компенсацией
func ExecuteTransaction(steps []BookingStep) error {
	for i, step := range steps {
		if err := step.Execute(); err != nil {
			fmt.Println("Error:", err)
			// Выполняем компенсацию для всех предыдущих шагов
			for j := i - 1; j >= 0; j-- {
				steps[j].Compensate()
			}
			return err
		}
	}
	return nil
}

func main() {
	steps := []BookingStep{
		&ReserveSeat{seatID: "A1"},
		&ProcessPayment{amount: 100.0},
	}

	if err := ExecuteTransaction(steps); err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction succeeded")
	}
}
