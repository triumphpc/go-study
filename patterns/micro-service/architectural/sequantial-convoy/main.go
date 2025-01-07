// Паттерн "Sequential Convoy" описывает проблему, которая возникает в многопоточных системах, когда несколько потоков
//или процессов ожидают доступа к ресурсу, но из-за последовательной обработки запросов производительность системы снижается.
//Это может произойти, когда один поток занимает ресурс на длительное время, заставляя другие потоки ожидать, что приводит к
//снижению общей пропускной способности системы.
//
//Когда возникает Sequential Convoy
//Блокировка ресурсов: Когда один поток или процесс занимает ресурс на длительное время, блокируя доступ для других.
//Последовательная обработка: Когда запросы обрабатываются последовательно, и один медленный запрос может задерживать другие.
//Ограниченные ресурсы: Когда ресурсы ограничены, и их доступность влияет на производительность системы.
//Высокая конкуренция: Когда множество потоков или процессов конкурируют за доступ к одному и тому же ресурсу.
//Проблемы, связанные с Sequential Convoy
//Снижение производительности: Общая пропускная способность системы снижается из-за ожидания потоков.
//Задержки: Время отклика системы увеличивается, так как потоки вынуждены ожидать освобождения ресурса.
//Неэффективное использование ресурсов: Ресурсы могут использоваться неэффективно, если они блокируются на длительное время.
//Увеличение времени ожидания: Потоки или процессы могут испытывать значительные задержки в ожидании доступа к ресурсу.
//Как избежать Sequential Convoy
//Увеличение параллелизма: Используйте параллельные структуры данных и алгоритмы, чтобы уменьшить блокировки.
//Оптимизация доступа к ресурсам: Минимизируйте время, в течение которого ресурсы заняты, и используйте более эффективные алгоритмы блокировки.
//Балансировка нагрузки: Распределяйте нагрузку равномерно между потоками или процессами, чтобы избежать перегрузки одного ресурса.
//Использование очередей: Применяйте очереди для управления доступом к ресурсам, чтобы избежать блокировок.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Message - структура, представляющая сообщение
type Message struct {
	Category string
	Content  string
}

// Queue - структура, представляющая очередь сообщений
type Queue struct {
	messages map[string][]Message
	mu       sync.Mutex
}

// NewQueue - создает новую очередь
func NewQueue() *Queue {
	return &Queue{messages: make(map[string][]Message)}
}

// Push - добавляет сообщение в очередь
func (q *Queue) Push(msg Message) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.messages[msg.Category] = append(q.messages[msg.Category], msg)
}

// Pull - извлекает сообщение из указанной категории
func (q *Queue) Pull(category string) (Message, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if msgs, exists := q.messages[category]; exists && len(msgs) > 0 {
		msg := msgs[0]
		q.messages[category] = msgs[1:]
		return msg, true
	}
	return Message{}, false
}

// Listener - структура, представляющая слушателя очереди
type Listener struct {
	ID       int
	Queue    *Queue
	Category string
}

// Listen - обрабатывает сообщения из указанной категории
func (l *Listener) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg, ok := l.Queue.Pull(l.Category)
		if !ok {
			break
		}
		fmt.Printf("Listener %d processing message from category '%s': %s\n", l.ID, msg.Category, msg.Content)
		time.Sleep(500 * time.Millisecond) // Симулируем обработку
	}
}

func main() {
	queue := NewQueue()

	// Добавляем сообщения в очередь
	queue.Push(Message{Category: "A", Content: "Message 1 in A"})
	queue.Push(Message{Category: "B", Content: "Message 1 in B"})
	queue.Push(Message{Category: "A", Content: "Message 2 in A"})
	queue.Push(Message{Category: "B", Content: "Message 2 in B"})

	var wg sync.WaitGroup

	// Создаем слушателей для каждой категории
	listeners := []Listener{
		{ID: 1, Queue: queue, Category: "A"},
		{ID: 2, Queue: queue, Category: "B"},
	}

	// Запускаем слушателей
	for _, listener := range listeners {
		wg.Add(1)
		go listener.Listen(&wg)
	}

	wg.Wait()
	fmt.Println("All messages processed")
}
