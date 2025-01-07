// Паттерн "Anti-Corruption Layer" (ACL) используется для изоляции и защиты вашей системы от внешних систем или сервисов,
//которые могут иметь несовместимые или нежелательные модели данных и поведения. Этот паттерн помогает минимизировать
//влияние изменений в одной системе на другую, обеспечивая стабильность и целостность вашей системы.

// Когда использовать Anti-Corruption Layer
//Интеграция с устаревшими системами: Когда ваша система должна взаимодействовать с устаревшими системами, которые имеют сложные или нестабильные интерфейсы.
//Несовместимые модели данных: Когда внешняя система использует модели данных, которые не соответствуют вашим внутренним моделям.
//Защита от изменений: Когда вы хотите защитить свою систему от изменений в API или поведении внешней системы.
//Сложная логика преобразования: Когда требуется сложная логика преобразования данных между системами.

// Плюсы Anti-Corruption Layer
//Изоляция изменений: Изменения в одной системе не влияют напрямую на другую, что упрощает управление изменениями.
//Упрощение интеграции: ACL упрощает интеграцию с внешними системами, предоставляя единый интерфейс для взаимодействия.
//Повышение надежности: Защищает вашу систему от ошибок и нестабильности внешних систем.
//Чистота кода: Сохраняет внутреннюю архитектуру вашей системы чистой и независимой от внешних систем.

// Минусы Anti-Corruption Layer
//Дополнительная сложность: ACL может потребовать дополнительной разработки и поддержки.
//Дополнительные затраты: ACL может потребовать дополнительных ресурсов для разработки и поддержки.
//Сложность отладки: ACL может сделать отладку сложнее, поскольку ошибки могут происходить в разных местах.

package main

import (
	"fmt"
	"strconv"
)

// Order - внутренняя модель заказа
type Order struct {
	ID    string
	Total float64
}

// LegacyOrder - модель заказа устаревшей системы
type LegacyOrder struct {
	OrderID string
	Amount  string
}

// OrderService - интерфейс для работы с заказами
type OrderService interface {
	GetOrder(id string) (*Order, error)
}

// LegacyOrderService - устаревшая система управления заказами
type LegacyOrderService struct{}

func (s *LegacyOrderService) FetchOrder(orderID string) (*LegacyOrder, error) {
	// Симуляция получения заказа из устаревшей системы
	return &LegacyOrder{OrderID: orderID, Amount: "100.50"}, nil
}

// AntiCorruptionLayer - слой защиты от устаревшей системы
type AntiCorruptionLayer struct {
	legacyService *LegacyOrderService
}

func NewAntiCorruptionLayer(legacyService *LegacyOrderService) *AntiCorruptionLayer {
	return &AntiCorruptionLayer{legacyService: legacyService}
}

func (acl *AntiCorruptionLayer) GetOrder(id string) (*Order, error) {
	legacyOrder, err := acl.legacyService.FetchOrder(id)
	if err != nil {
		return nil, err
	}

	// Преобразование данных из устаревшей системы в внутреннюю модель
	total, err := parseAmount(legacyOrder.Amount)
	if err != nil {
		return nil, err
	}

	return &Order{ID: legacyOrder.OrderID, Total: total}, nil
}

func parseAmount(amount string) (float64, error) {
	return strconv.ParseFloat(amount, 64)
}

func main() {
	legacyService := &LegacyOrderService{}
	acl := NewAntiCorruptionLayer(legacyService)

	order, err := acl.GetOrder("123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Order ID: %s, Total: %.2f\n", order.ID, order.Total)
}
