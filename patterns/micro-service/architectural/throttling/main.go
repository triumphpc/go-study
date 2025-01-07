// Паттерн "Throttling" используется для управления и ограничения количества запросов или операций, которые могут
//быть выполнены в системе за определенный период времени. Это помогает предотвратить перегрузку системы, обеспечивая
//стабильную производительность и защищая ресурсы от чрезмерного использования. Паттерн особенно полезен в высоконагруженных
//системах и при взаимодействии с внешними сервисами.
//
//Когда использовать Throttling Pattern
//Защита ресурсов: Когда необходимо защитить ресурсы от перегрузки и чрезмерного использования.
//Управление нагрузкой: Когда требуется контролировать количество запросов или операций, чтобы обеспечить стабильную производительность.
//Интеграция с внешними сервисами: Когда взаимодействие с внешними сервисами ограничено по количеству запросов.
//Обеспечение качества обслуживания: Когда необходимо гарантировать, что система может обрабатывать запросы в пределах своих возможностей.

//Плюсы Throttling Pattern
//Стабильная производительность: Обеспечивает стабильную производительность системы, предотвращая перегрузку.
//Защита от атак: Помогает защитить систему от атак, таких как DDoS, ограничивая количество запросов.
//Управление ресурсами: Позволяет эффективно управлять ресурсами, распределяя нагрузку равномерно.
//Гибкость: Легко настраивается в зависимости от требований системы и доступных ресурсов.

//Минусы Throttling Pattern
//Задержки: Может добавить задержки в обработку запросов, если они превышают установленный лимит.
//Ограничение доступности: Может ограничить доступность системы для пользователей, если лимиты слишком жесткие.
//Сложность настройки: Требует тщательной настройки параметров ограничения, чтобы избежать негативного влияния на пользователей.
//Потенциальные блокировки: Неправильная настройка может привести к блокировке легитимных запросов.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Throttler - структура для управления ограничением запросов
type Throttler struct {
	rate      int           // Количество запросов в секунду
	interval  time.Duration // Интервал времени для ограничения
	tokens    chan struct{} // Канал для токенов
	stopChan  chan struct{} // Канал для остановки
	waitGroup sync.WaitGroup
}

// NewThrottler - создает новый ограничитель
func NewThrottler(rate int) *Throttler {
	t := &Throttler{
		rate:     rate,
		interval: time.Second / time.Duration(rate),
		tokens:   make(chan struct{}, rate),
		stopChan: make(chan struct{}),
	}

	// Запускаем горутину для пополнения токенов
	t.waitGroup.Add(1)
	go t.start()

	return t
}

// start - пополняет токены в канале
func (t *Throttler) start() {
	defer t.waitGroup.Done()
	ticker := time.NewTicker(t.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case t.tokens <- struct{}{}:
			default:
			}
		case <-t.stopChan:
			return
		}
	}
}

// Allow - проверяет, можно ли выполнить запрос
func (t *Throttler) Allow() bool {
	select {
	case <-t.tokens:
		return true
	default:
		return false
	}
}

// Stop - останавливает ограничитель
func (t *Throttler) Stop() {
	close(t.stopChan)
	t.waitGroup.Wait()
}

func main() {
	throttler := NewThrottler(5) // Ограничиваем до 5 запросов в секунду
	defer throttler.Stop()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !throttler.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		fmt.Fprintln(w, "Request processed")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
