// Паттерн "Geode" используется для управления данными в распределенных системах, обеспечивая высокую доступность и
//отказоустойчивость. Этот паттерн предполагает распределение данных по множеству узлов или географически распределенных
//центров обработки данных, что позволяет системе оставаться доступной даже в случае сбоя одного или нескольких узлов.
//Geode-паттерн часто используется в системах, где требуется высокая доступность и минимальные задержки при доступе к данным.
//
//Когда использовать Geode Pattern
//Высокая доступность: Когда необходимо обеспечить доступность данных даже в случае сбоя одного или нескольких узлов.
//Географическое распределение: Когда данные должны быть доступны в разных географических регионах для улучшения производительности и доступности.
//Отказоустойчивость: Когда требуется обеспечить устойчивость к сбоям и минимизировать время простоя.
//Масштабируемость: Когда необходимо масштабировать систему для обработки увеличивающегося объема данных и нагрузки.
//Плюсы Geode Pattern
//Высокая доступность: Обеспечивает доступность данных даже в случае сбоя одного или нескольких узлов.
//Отказоустойчивость: Повышает устойчивость системы к сбоям и минимизирует время простоя.
//Масштабируемость: Позволяет легко масштабировать систему, добавляя новые узлы или центры обработки данных.
//Улучшение производительности: Распределение данных по множеству узлов может улучшить производительность и уменьшить задержки.
//Минусы Geode Pattern
//Сложность управления: Требует управления распределением данных и координацией между узлами.
//Согласованность данных: Необходимо обеспечить согласованность данных между различными узлами или центрами обработки данных.
//Затраты на инфраструктуру: Может потребовать дополнительных ресурсов и инфраструктуры для поддержки распределенной системы.
//Сложность реализации: Требует тщательной реализации и управления, чтобы обеспечить надежность и производительность.

package main

import (
	"fmt"
	"sync"
)

// Node - структура, представляющая узел в распределенной системе
type Node struct {
	ID   int
	Data map[string]string
	mu   sync.RWMutex
}

// NewNode - создает новый узел
func NewNode(id int) *Node {
	return &Node{
		ID:   id,
		Data: make(map[string]string),
	}
}

// Set - сохраняет данные на узле
func (n *Node) Set(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Data[key] = value
}

// Get - извлекает данные с узла
func (n *Node) Get(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	value, exists := n.Data[key]
	return value, exists
}

// DistributedSystem - структура для управления распределенной системой
type DistributedSystem struct {
	nodes []*Node
}

// NewDistributedSystem - создает новую распределенную систему
func NewDistributedSystem(numNodes int) *DistributedSystem {
	nodes := make([]*Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = NewNode(i)
	}
	return &DistributedSystem{nodes: nodes}
}

// Set - сохраняет данные в распределенной системе
func (ds *DistributedSystem) Set(key, value string) {
	node := ds.getNodeForKey(key)
	node.Set(key, value)
}

// Get - извлекает данные из распределенной системы
func (ds *DistributedSystem) Get(key string) (string, bool) {
	node := ds.getNodeForKey(key)
	return node.Get(key)
}

// getNodeForKey - определяет, на каком узле хранить данные
func (ds *DistributedSystem) getNodeForKey(key string) *Node {
	hash := int(key[0]) % len(ds.nodes)
	return ds.nodes[hash]
}

func main() {
	ds := NewDistributedSystem(3)

	// Сохраняем данные
	ds.Set("key1", "value1")
	ds.Set("key2", "value2")
	ds.Set("key3", "value3")

	// Извлекаем данные
	if value, exists := ds.Get("key1"); exists {
		fmt.Printf("Retrieved value for 'key1': %s\n", value)
	} else {
		fmt.Println("Key 'key1' not found")
	}
}
