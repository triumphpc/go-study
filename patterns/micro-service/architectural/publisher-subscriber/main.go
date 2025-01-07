// Паттерн "Publisher-Subscriber" (или "Pub-Sub") — это архитектурный шаблон, который используется для асинхронного
//обмена сообщениями между компонентами системы. В этом паттерне издатели (publishers) отправляют сообщения, не зная,
//кто их получит, а подписчики (subscribers) получают сообщения, на которые они подписаны, не зная, кто их отправил.
//Это позволяет создать гибкую и масштабируемую систему, где компоненты слабо связаны между собой.
//
//Когда использовать Publisher-Subscriber Pattern
//Асинхронная коммуникация: Когда необходимо обеспечить асинхронное взаимодействие между компонентами системы.
//Слабо связанная архитектура: Когда требуется уменьшить зависимость между компонентами, чтобы улучшить масштабируемость и гибкость.
//Масштабируемость: Когда необходимо легко добавлять или удалять компоненты без изменения других частей системы.
//Обработка событий: Когда система должна реагировать на события, такие как изменения состояния или обновления данных.
//Плюсы Publisher-Subscriber Pattern
//Слабая связанность: Компоненты не зависят друг от друга, что упрощает их замену и обновление.
//Масштабируемость: Легко добавлять новых подписчиков или издателей без изменения существующих компонентов.
//Гибкость: Позволяет легко изменять и расширять функциональность системы.
//Асинхронность: Обеспечивает асинхронное взаимодействие, что может улучшить производительность и отзывчивость системы.
//Минусы Publisher-Subscriber Pattern
//Сложность отладки: Асинхронное взаимодействие может усложнить отладку и мониторинг системы.
//Задержки: Может добавить задержки в доставку сообщений, особенно в распределенных системах.
//Потеря сообщений: В некоторых реализациях возможна потеря сообщений, если подписчики не успевают их обработать.
//Сложность управления: Требует управления подписками и маршрутизацией сообщений.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Publisher - структура, представляющая издателя
type Publisher struct {
	subscribers map[chan string]struct{}
	mu          sync.RWMutex
}

// NewPublisher - создает нового издателя
func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[chan string]struct{}),
	}
}

// Subscribe - добавляет нового подписчика
func (p *Publisher) Subscribe() chan string {
	ch := make(chan string)
	p.mu.Lock()
	p.subscribers[ch] = struct{}{}
	p.mu.Unlock()
	return ch
}

// Unsubscribe - удаляет подписчика
func (p *Publisher) Unsubscribe(ch chan string) {
	p.mu.Lock()
	delete(p.subscribers, ch)
	close(ch)
	p.mu.Unlock()
}

// Publish - отправляет сообщение всем подписчикам
func (p *Publisher) Publish(msg string) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for ch := range p.subscribers {
		ch <- msg
	}
}

func main() {
	publisher := NewPublisher()

	// Создаем подписчиков
	sub1 := publisher.Subscribe()
	sub2 := publisher.Subscribe()

	// Запускаем горутины для обработки сообщений подписчиками
	go func() {
		for msg := range sub1 {
			fmt.Printf("Subscriber 1 received: %s\n", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Printf("Subscriber 2 received: %s\n", msg)
		}
	}()

	// Публикуем сообщения
	publisher.Publish("Hello, World!")
	time.Sleep(1 * time.Second)
	publisher.Publish("Another message")

	// Отписываем подписчиков
	publisher.Unsubscribe(sub1)
	publisher.Unsubscribe(sub2)

	time.Sleep(1 * time.Second)
}
