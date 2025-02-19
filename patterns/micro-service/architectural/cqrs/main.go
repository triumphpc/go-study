// Паттерн "Command and Query Responsibility Segregation" (CQRS) разделяет операции чтения и записи в системе, позволяя
// им использовать разные модели данных. Это позволяет оптимизировать каждую операцию в зависимости от её специфики,
//улучшая производительность и масштабируемость системы.

// Когда использовать CQRS
//Сложные доменные модели: Когда у вас есть сложные бизнес-правила и модели данных, которые трудно поддерживать в одной модели.
//Высокая нагрузка на чтение: Когда система испытывает большую нагрузку на операции чтения, и их нужно оптимизировать отдельно от операций записи.
//Требования к масштабируемости: Когда необходимо масштабировать операции чтения и записи независимо друг от друга.
//Разные требования к безопасности: Когда операции чтения и записи имеют разные требования к безопасности и доступу.

//Плюсы CQRS
//Оптимизация производительности: Позволяет оптимизировать операции чтения и записи независимо друг от друга.
//Масштабируемость: Обеспечивает возможность масштабирования операций чтения и записи отдельно.
//Упрощение модели: Разделение моделей данных упрощает управление сложными доменными моделями.
//Гибкость: Позволяет использовать разные технологии и подходы для чтения и записи.

//Минусы CQRS
//Сложность реализации: Увеличивает сложность системы из-за необходимости управления двумя моделями данных.
//Согласованность данных: Может привести к временной несогласованности данных между моделями чтения и записи.
//Затраты на поддержку: Требует дополнительных усилий для поддержки и управления двумя моделями данных.

package main

import (
	"fmt"
)

// OrderCommand - модель для операций записи
type OrderCommand struct {
	ID    string
	Items []string
}

// OrderQuery - модель для операций чтения
type OrderQuery struct {
	ID    string
	Total int
}

// OrderService - сервис для управления заказами
type OrderService struct {
	orders map[string]OrderCommand
}

func NewOrderService() *OrderService {
	return &OrderService{orders: make(map[string]OrderCommand)}
}

// CreateOrder - операция записи
func (s *OrderService) CreateOrder(id string, items []string) {
	s.orders[id] = OrderCommand{ID: id, Items: items}
	fmt.Println("Order created:", id)
}

// GetOrderTotal - операция чтения
func (s *OrderService) GetOrderTotal(id string) int {
	order, exists := s.orders[id]
	if !exists {
		fmt.Println("Order not found:", id)
		return 0
	}
	// Симуляция расчета общей стоимости
	total := len(order.Items) * 10
	fmt.Println("Order total for", id, "is", total)
	return total
}

func main() {
	service := NewOrderService()

	// Создаем заказ
	service.CreateOrder("1", []string{"item1", "item2", "item3"})

	// Получаем общую стоимость заказа
	service.GetOrderTotal("1")
}
