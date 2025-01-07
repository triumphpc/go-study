// Паттерн "Choreography" используется в распределенных системах для координации взаимодействия между несколькими
//сервисами без центрального управляющего компонента. В отличие от паттерна "Orchestration", где один компонент
//управляет всеми взаимодействиями, в хореографии каждый сервис знает, как реагировать на события и взаимодействовать
//с другими сервисами. Это позволяет создать более гибкую и масштабируемую архитектуру.
//
//Когда использовать Choreography Pattern
//Микросервисные архитектуры: Когда система состоит из множества независимых сервисов, которые должны взаимодействовать друг с другом.
//Гибкость и масштабируемость: Когда требуется гибкость в добавлении новых сервисов и масштабировании системы.
//Событийно-ориентированные системы: Когда взаимодействие между сервисами основано на событиях.
//Избежание единой точки отказа: Когда необходимо избежать зависимости от центрального управляющего компонента.

//Плюсы Choreography Pattern
//Гибкость: Легко добавлять новые сервисы и изменять существующие без изменения центрального компонента.
//Масштабируемость: Каждый сервис может быть масштабирован независимо, что упрощает управление нагрузкой.
//Отказоустойчивость: Отсутствие центрального управляющего компонента снижает риск единой точки отказа.
//Упрощение управления: Каждый сервис управляет своим собственным состоянием и взаимодействиями.

//Минусы Choreography Pattern
//Сложность координации: Может быть сложно координировать взаимодействие между множеством сервисов.
//Трудности отладки: Отладка распределенной системы без центрального управляющего компонента может быть сложной.
//Потенциальные задержки: Взаимодействие между сервисами может привести к задержкам, особенно если они зависят от сети.
//Сложность управления событиями: Необходимо управлять событиями и их обработкой, чтобы избежать несогласованности данных.

package main

import (
	"fmt"
	"time"
)

// Event - структура, представляющая событие
type Event struct {
	Name string
	Data string
}

// Service - интерфейс для сервисов
type Service interface {
	HandleEvent(event Event)
}

// OrderService - сервис для обработки заказов
type OrderService struct{}

func (s *OrderService) HandleEvent(event Event) {
	if event.Name == "OrderCreated" {
		fmt.Println("OrderService: Processing order", event.Data)
		// Генерируем событие для InventoryService
		inventoryService.HandleEvent(Event{Name: "CheckInventory", Data: event.Data})
	}
}

// InventoryService - сервис для управления запасами
type InventoryService struct{}

func (s *InventoryService) HandleEvent(event Event) {
	if event.Name == "CheckInventory" {
		fmt.Println("InventoryService: Checking inventory for", event.Data)
		// Генерируем событие для ShippingService
		shippingService.HandleEvent(Event{Name: "PrepareShipment", Data: event.Data})
	}
}

// ShippingService - сервис для управления доставкой
type ShippingService struct{}

func (s *ShippingService) HandleEvent(event Event) {
	if event.Name == "PrepareShipment" {
		fmt.Println("ShippingService: Preparing shipment for", event.Data)
	}
}

var (
	orderService     = &OrderService{}
	inventoryService = &InventoryService{}
	shippingService  = &ShippingService{}
)

func main() {
	// Симулируем создание заказа
	orderService.HandleEvent(Event{Name: "OrderCreated", Data: "Order123"})

	// Ждем завершения всех операций
	time.Sleep(1 * time.Second)
}
