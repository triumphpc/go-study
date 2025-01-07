// Паттерн "Gateway Aggregation" используется для объединения и координации запросов к нескольким внутренним сервисам
//через единый шлюз (gateway). Этот паттерн позволяет клиентам взаимодействовать с различными сервисами через один интерфейс,
//упрощая архитектуру и улучшая производительность за счет уменьшения количества сетевых вызовов.
//
//Когда использовать Gateway Aggregation Pattern
//Микросервисные архитектуры: Когда система состоит из множества микросервисов, и необходимо предоставить единый интерфейс для взаимодействия с ними.
//Оптимизация производительности: Когда требуется уменьшить количество сетевых вызовов, объединяя их в один запрос.
//Упрощение клиентской логики: Когда необходимо упростить клиентскую логику, предоставляя единый API для доступа к нескольким сервисам.
//Централизованное управление: Когда требуется централизованное управление доступом и безопасностью для нескольких сервисов.

//Плюсы Gateway Aggregation Pattern
//Упрощение клиентской логики: Клиенты взаимодействуют с одним API, что упрощает их реализацию и поддержку.
//Улучшение производительности: Объединение нескольких запросов в один уменьшает количество сетевых вызовов и задержек.
//Централизованное управление: Позволяет централизованно управлять доступом, безопасностью и мониторингом.
//Гибкость: Легко адаптируется к изменениям в архитектуре микросервисов.

//Минусы Gateway Aggregation Pattern
//Сложность реализации: Требует тщательной реализации логики агрегации и управления запросами.
//Единая точка отказа: Шлюз может стать единой точкой отказа, если не предусмотрены механизмы отказоустойчивости.
//Задержки: Может добавить задержки в обработку запросов из-за агрегации и координации.
//Зависимость от шлюза: Все взаимодействия зависят от шлюза, что может ограничить гибкость системы.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// ServiceResponse - структура для хранения ответа от сервиса
type ServiceResponse struct {
	ServiceName string `json:"service_name"`
	Data        string `json:"data"`
}

// fetchData - функция для получения данных от сервиса
func fetchData(serviceName, url string, wg *sync.WaitGroup, responses chan<- ServiceResponse) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from %s: %v\n", serviceName, err)
		return
	}
	defer resp.Body.Close()

	var data string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding response from %s: %v\n", serviceName, err)
		return
	}

	responses <- ServiceResponse{ServiceName: serviceName, Data: data}
}

// gatewayHandler - обработчик для агрегации данных
func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	responses := make(chan ServiceResponse, 2)

	// Запрашиваем данные от двух внутренних сервисов
	wg.Add(2)
	go fetchData("ServiceA", "http://service-a.local/data", &wg, responses)
	go fetchData("ServiceB", "http://service-b.local/data", &wg, responses)

	wg.Wait()
	close(responses)

	// Агрегируем результаты
	var aggregatedResponses []ServiceResponse
	for response := range responses {
		aggregatedResponses = append(aggregatedResponses, response)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggregatedResponses)
}

func main() {
	http.HandleFunc("/aggregate", gatewayHandler)
	fmt.Println("Gateway server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
