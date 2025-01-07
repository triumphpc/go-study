// Паттерн "Gateway Offloading" используется для переноса определенных задач и функций с серверов приложений на шлюз (gateway).
//Это позволяет разгрузить серверы приложений, улучшить производительность и упростить архитектуру. Шлюз может обрабатывать
//такие задачи, как аутентификация, шифрование, кэширование и сжатие, освобождая серверы приложений для выполнения основной бизнес-логики.
//
//Когда использовать Gateway Offloading Pattern
//Разгрузка серверов приложений: Когда необходимо уменьшить нагрузку на серверы приложений, передав часть задач на шлюз.
//Улучшение производительности: Когда требуется улучшить производительность системы за счет оптимизации обработки запросов.
//Централизованное управление: Когда необходимо централизованно управлять задачами, такими как аутентификация и шифрование.
//Упрощение архитектуры: Когда требуется упростить архитектуру приложения, переместив общие функции на шлюз.

//Плюсы Gateway Offloading Pattern
//Улучшение производительности: Разгружает серверы приложений, позволяя им сосредоточиться на выполнении основной бизнес-логики.
//Централизованное управление: Обеспечивает централизованное управление задачами, такими как аутентификация и шифрование.
//Гибкость: Легко адаптируется к изменениям в архитектуре и требованиям системы.
//Упрощение архитектуры: Упрощает архитектуру приложения, перемещая общие функции на шлюз.

//Минусы Gateway Offloading Pattern
//Единая точка отказа: Шлюз может стать единой точкой отказа, если не предусмотрены механизмы отказоустойчивости.
//Сложность настройки: Требует тщательной настройки и управления задачами, переданными на шлюз.
//Задержки: Может добавить задержки в обработку запросов из-за дополнительного уровня обработки.
//Зависимость от шлюза: Все взаимодействия зависят от шлюза, что может ограничить гибкость системы.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Cache - структура для хранения кэша
type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewCache - создает новый кэш
func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

// Get - получает данные из кэша
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

// Set - сохраняет данные в кэш
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// authenticate - функция для аутентификации запросов
func authenticate(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	return token == "valid-token"
}

// gatewayHandler - обработчик для обработки запросов
func gatewayHandler(cache *Cache, w http.ResponseWriter, r *http.Request) {
	if !authenticate(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	key := r.URL.Path
	if value, exists := cache.Get(key); exists {
		fmt.Fprintf(w, "Cached response: %s", value)
		return
	}

	// Симулируем обработку запроса
	time.Sleep(500 * time.Millisecond)
	response := "Processed response"

	cache.Set(key, response)
	fmt.Fprintf(w, response)
}

func main() {
	cache := NewCache()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gatewayHandler(cache, w, r)
	})

	fmt.Println("Gateway server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
