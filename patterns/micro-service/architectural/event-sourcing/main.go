// Паттерн "Event Sourcing" — это архитектурный шаблон, который сохраняет состояние системы как последовательность событий.
//Вместо хранения текущего состояния объекта, система сохраняет все изменения, которые привели к этому состоянию. Это позволяет
//легко восстанавливать прошлые состояния и обеспечивает полную историю изменений.

// Когда использовать Event Sourcing
//Требования к аудиту: Когда необходимо иметь полную историю изменений для целей аудита или соответствия нормативным требованиям.
//Сложные бизнес-правила: Когда бизнес-логика сложна и требует отслеживания всех изменений.
//Восстановление состояния: Когда необходимо восстанавливать прошлые состояния системы для анализа или отладки.
//Интеграция с CQRS: Часто используется вместе с CQRS для разделения операций чтения и записи.

//Плюсы Event Sourcing
//Полная история изменений: Обеспечивает полную историю всех изменений, что полезно для аудита и анализа.
//Восстановление состояния: Позволяет восстанавливать любое прошлое состояние системы.
//Гибкость: Позволяет легко добавлять новые функции, такие как уведомления или аналитика, на основе событий.
//Согласованность: Обеспечивает согласованность данных через события.

//Минусы Event Sourcing
//Сложность реализации: Увеличивает сложность системы из-за необходимости управления событиями и их последовательностью.
//Объем данных: Может привести к значительному увеличению объема данных, так как все изменения сохраняются.
//Сложность чтения: Получение текущего состояния может быть сложным, так как требует воспроизведения всех событий.
//Зависимость от событий: Изменение структуры событий может быть сложным и требует тщательного управления версиями.

package main

import (
	"fmt"
)

// Event - интерфейс для событий
type Event interface {
	Apply(account *Account)
}

// AccountCreated - событие создания счета
type AccountCreated struct {
	ID string
}

func (e AccountCreated) Apply(account *Account) {
	account.ID = e.ID
}

// MoneyDeposited - событие депозита средств
type MoneyDeposited struct {
	Amount float64
}

func (e MoneyDeposited) Apply(account *Account) {
	account.Balance += e.Amount
}

// Account - структура счета
type Account struct {
	ID      string
	Balance float64
	events  []Event
}

// NewAccount - создает новый счет
func NewAccount(id string) *Account {
	account := &Account{}
	account.ApplyEvent(AccountCreated{ID: id})
	return account
}

// ApplyEvent - применяет событие к счету
func (a *Account) ApplyEvent(event Event) {
	event.Apply(a)
	a.events = append(a.events, event)
}

// Deposit - добавляет средства на счет
func (a *Account) Deposit(amount float64) {
	a.ApplyEvent(MoneyDeposited{Amount: amount})
}

// GetBalance - возвращает текущий баланс
func (a *Account) GetBalance() float64 {
	return a.Balance
}

func main() {
	account := NewAccount("12345")
	account.Deposit(100.0)
	account.Deposit(50.0)

	fmt.Printf("Account ID: %s, Balance: %.2f\n", account.ID, account.GetBalance())
}
