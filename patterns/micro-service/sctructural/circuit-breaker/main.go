// Паттерн "Circuit Breaker" (предохранитель) используется для повышения устойчивости и надежности распределенных систем.
//Он помогает предотвратить каскадные сбои в системе, ограничивая количество неудачных попыток доступа к внешним сервисам
//или ресурсам. Этот паттерн особенно полезен, когда вы взаимодействуете с ненадежными или перегруженными сервисами.
//
//Когда использовать Circuit Breaker
//Ненадежные внешние сервисы: Если вы взаимодействуете с сервисами, которые могут быть недоступны или медленно отвечать.
//Ограничение ресурсов: Чтобы избежать перегрузки системы из-за большого количества неудачных попыток.
//Улучшение пользовательского опыта: Быстрое возвращение ошибки вместо долгого ожидания ответа от сервиса.

//Как работает Circuit Breaker
//Closed (Закрытое) состояние: Все запросы проходят. Если количество ошибок превышает порог, переключается в Open состояние.
//Open (Открытое) состояние: Запросы блокируются и сразу возвращается ошибка. Через определенное время переключается в Half-Open.
//Half-Open (Полуоткрытое) состояние: Разрешает ограниченное количество запросов для проверки, восстановился ли сервис. Если успешные,
//переключается в Closed, если нет — обратно в Open.

package main

import (
	"errors"
	"fmt"
	"time"
)

type CircuitBreaker struct {
	failures     int
	maxFailures  int
	state        string
	lastFailure  time.Time
	retryTimeout time.Duration
}

func NewCircuitBreaker(maxFailures int, retryTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		retryTimeout: retryTimeout,
		state:        "CLOSED",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	switch cb.state {
	case "OPEN":
		if time.Since(cb.lastFailure) > cb.retryTimeout {
			cb.state = "HALF-OPEN"
		} else {
			return errors.New("circuit breaker is open")
		}
	case "HALF-OPEN":
		// Allow a single request to test the service
	}

	err := fn()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.maxFailures {
			cb.state = "OPEN"
		}
		return err
	}

	cb.reset()
	return nil
}

func (cb *CircuitBreaker) reset() {
	cb.failures = 0
	cb.state = "CLOSED"
}

func main() {
	cb := NewCircuitBreaker(3, 2*time.Second)

	for i := 0; i < 10; i++ {
		err := cb.Call(func() error {
			// Simulate a failing service
			return errors.New("service failure")
		})

		if err != nil {
			fmt.Printf("Attempt %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Attempt %d: success\n", i+1)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
