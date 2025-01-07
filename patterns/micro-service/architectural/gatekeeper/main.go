// Паттерн "Gatekeeper" используется для управления доступом к ресурсам или сервисам, обеспечивая контроль и безопасность. Он действует как посредник между клиентами и ресурсами, проверяя запросы и применяя политики доступа. Этот паттерн часто используется для защиты ресурсов от несанкционированного доступа и для обеспечения соблюдения бизнес-правил.
//
//Когда использовать Gatekeeper Pattern
//Контроль доступа: Когда необходимо ограничить доступ к ресурсам или сервисам на основе определенных правил или политик.
//Безопасность: Когда требуется защита ресурсов от несанкционированного доступа или атак.
//Логирование и мониторинг: Когда необходимо вести учет и мониторинг доступа к ресурсам.
//Управление нагрузкой: Когда требуется ограничить количество запросов к ресурсу для предотвращения перегрузки.

//Плюсы Gatekeeper Pattern
//Улучшенная безопасность: Обеспечивает защиту ресурсов от несанкционированного доступа и атак.
//Централизованный контроль: Позволяет централизованно управлять доступом и политиками безопасности.
//Гибкость: Легко адаптируется к изменяющимся требованиям безопасности и бизнес-правилам.
//Мониторинг и аудит: Обеспечивает возможность логирования и мониторинга доступа к ресурсам.

//Минусы Gatekeeper Pattern
//Единая точка отказа: Если gatekeeper выходит из строя, это может привести к недоступности ресурсов.
//Задержки: Может добавить задержки в обработку запросов из-за дополнительного уровня проверки.
//Сложность настройки: Требует тщательной настройки и управления политиками доступа.
//Зависимость от конфигурации: Изменения в политике доступа могут потребовать обновления конфигурации gatekeeper.

package main

import (
	"fmt"
	"net/http"
)

// Gatekeeper - структура для управления доступом
type Gatekeeper struct {
	validAPIKeys map[string]bool
}

// NewGatekeeper - создает новый экземпляр Gatekeeper
func NewGatekeeper(validKeys []string) *Gatekeeper {
	keysMap := make(map[string]bool)
	for _, key := range validKeys {
		keysMap[key] = true
	}
	return &Gatekeeper{validAPIKeys: keysMap}
}

// Middleware - промежуточный обработчик для проверки доступа
func (g *Gatekeeper) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if !g.validAPIKeys[apiKey] {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	validKeys := []string{"key1", "key2", "key3"}
	gatekeeper := NewGatekeeper(validKeys)

	http.Handle("/", gatekeeper.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Access granted")
	})))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
