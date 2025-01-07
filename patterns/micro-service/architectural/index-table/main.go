// Паттерн "Index Table" используется для оптимизации поиска и доступа к данным в базе данных. Он предполагает создание
//дополнительной таблицы индексов, которая хранит ссылки на данные в основной таблице. Это позволяет ускорить операции
//поиска и сортировки, особенно в больших наборах данных.
//
//Когда использовать Index Table Pattern
//Ускорение поиска: Когда необходимо ускорить операции поиска и сортировки в больших наборах данных.
//Оптимизация запросов: Когда запросы к базе данных становятся медленными из-за большого объема данных.
//Частые операции выборки: Когда приложение часто выполняет операции выборки по определенным полям.
//Сложные запросы: Когда запросы содержат сложные условия фильтрации и сортировки.

//Плюсы Index Table Pattern
//Улучшение производительности: Значительно ускоряет операции поиска и сортировки.
//Оптимизация запросов: Позволяет оптимизировать сложные запросы, уменьшая время их выполнения.
//Гибкость: Позволяет создавать индексы для различных полей и комбинаций полей.
//Снижение нагрузки: Уменьшает нагрузку на базу данных за счет более эффективного выполнения запросов.

// Минусы Index Table Pattern
// Дополнительное пространство: Требует дополнительного пространства для хранения индексов.
// Сложность управления: Увеличивает сложность управления базой данных из-за необходимости обновления индексов при изменении данных.
// Задержки при записи: Может увеличить время выполнения операций вставки, обновления и удаления из-за необходимости обновления индексов.
// Избыточность: Может привести к избыточности данных, если индексы не используются эффективно.
package main

import (
	"fmt"
)

// Record - структура для хранения данных
type Record struct {
	ID    int
	Name  string
	Email string
}

// IndexTable - структура для хранения индексов
type IndexTable struct {
	byName  map[string][]*Record
	byEmail map[string][]*Record
}

// NewIndexTable - создает новую таблицу индексов
func NewIndexTable(records []*Record) *IndexTable {
	index := &IndexTable{
		byName:  make(map[string][]*Record),
		byEmail: make(map[string][]*Record),
	}

	for _, record := range records {
		index.byName[record.Name] = append(index.byName[record.Name], record)
		index.byEmail[record.Email] = append(index.byEmail[record.Email], record)
	}

	return index
}

// FindByName - ищет записи по имени
func (index *IndexTable) FindByName(name string) []*Record {
	return index.byName[name]
}

// FindByEmail - ищет записи по email
func (index *IndexTable) FindByEmail(email string) []*Record {
	return index.byEmail[email]
}

func main() {
	records := []*Record{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Alice", Email: "alice2@example.com"},
	}

	indexTable := NewIndexTable(records)

	// Поиск по имени
	results := indexTable.FindByName("Alice")
	fmt.Println("Records with name 'Alice':")
	for _, record := range results {
		fmt.Printf("ID: %d, Email: %s\n", record.ID, record.Email)
	}

	// Поиск по email
	results = indexTable.FindByEmail("bob@example.com")
	fmt.Println("Records with email 'bob@example.com':")
	for _, record := range results {
		fmt.Printf("ID: %d, Name: %s\n", record.ID, record.Name)
	}
}
