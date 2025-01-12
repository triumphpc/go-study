// Gob (Go Object Binary) - это встроенный в Go формат сериализации данных, разработанный специально для Go.Преимущества gob:
//
// Оптимизирован для Go (работает быстрее JSON для Go структур)
// Поддерживает все типы Go (включая интерфейсы, каналы)
// Компактнее JSON
// Типобезопасный
// Автоматическая обработка указателей и циклических ссылок
// Эффективен для передачи данных между Go программами
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

// EventKind - тип события
type EventKind string

const (
	AddKind      EventKind = "add"
	CheckoutKind EventKind = "checkout"
)

// Event - интерфейс для всех событий
type Event interface {
	Kind() EventKind
}

// Add - событие добавления товара
type Add struct {
	ProductID string
	Quantity  int
	Price     float64
}

func (a *Add) Kind() EventKind {
	return AddKind
}

// Checkout - событие оформления заказа
type Checkout struct {
	OrderID string
	Total   float64
}

func (c *Checkout) Kind() EventKind {
	return CheckoutKind
}

// Обработчики событий
func handleAdd(a *Add) {
	fmt.Printf("Handling Add event: Product=%s, Quantity=%d, Price=%.2f\n",
		a.ProductID, a.Quantity, a.Price)
}

func handleCheckout(c *Checkout) {
	fmt.Printf("Handling Checkout event: Order=%s, Total=%.2f\n",
		c.OrderID, c.Total)
}

// Основной обработчик событий
func eventHandler(r io.Reader) error {
	// Создаем новый декодер gob
	dec := gob.NewDecoder(r)

	// Бесконечный цикл чтения событий
	for {
		var e Event
		// Декодируем следующее событие
		err := dec.Decode(&e)
		if err == io.EOF {
			// Достигнут конец потока
			break
		}
		if err != nil {
			return err
		}

		// Обработка события в зависимости от его типа
		switch e.Kind() {
		case AddKind:
			handleAdd(e.(*Add))
		case CheckoutKind:
			handleCheckout(e.(*Checkout))
		default:
			return fmt.Errorf("unknown event kind: %s", e.Kind())
		}
	}
	return nil
}

func main() {
	// Регистрируем типы для gob
	gob.Register(&Add{})
	gob.Register(&Checkout{})

	// Создаем буфер для тестирования
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	// Создаем несколько тестовых событий
	events := []Event{
		&Add{
			ProductID: "123",
			Quantity:  2,
			Price:     29.99,
		},
		&Checkout{
			OrderID: "ORDER-001",
			Total:   59.98,
		},
	}

	// Кодируем события
	for _, e := range events {
		if err := enc.Encode(&e); err != nil {
			log.Fatal(err)
		}
	}

	// Обрабатываем события
	if err := eventHandler(&buf); err != nil {
		log.Fatal(err)
	}
}
