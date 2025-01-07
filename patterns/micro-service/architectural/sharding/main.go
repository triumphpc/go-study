// Паттерн "Sharding" используется для горизонтального разделения данных на более мелкие, управляемые части,
//называемые "шардами". Каждый шард хранится и обрабатывается независимо, что позволяет распределять нагрузку и
//улучшать производительность системы. Этот паттерн особенно полезен для работы с большими объемами данных и высоконагруженными системами.
//
//Когда использовать Sharding Pattern
//Большие объемы данных: Когда объем данных превышает возможности одного сервера или базы данных.
//Высокая нагрузка: Когда система испытывает высокую нагрузку, и необходимо распределить ее между несколькими серверами.
//Масштабируемость: Когда требуется горизонтальное масштабирование для обработки увеличивающегося объема данных и нагрузки.
//Улучшение производительности: Когда необходимо уменьшить время отклика и повысить производительность системы.

//Плюсы Sharding Pattern
//Масштабируемость: Позволяет горизонтально масштабировать систему, добавляя новые шарды по мере необходимости.
//Улучшение производительности: Уменьшает нагрузку на каждый отдельный сервер, что может улучшить время отклика и общую производительность.
//Распределение нагрузки: Распределяет нагрузку между несколькими серверами, что может повысить надежность и отказоустойчивость.
//Гибкость: Позволяет использовать разные стратегии шардирования в зависимости от требований системы.

// Минусы Sharding Pattern
// Сложность реализации: Требует тщательного проектирования и управления, чтобы обеспечить согласованность и целостность данных.
// Сложность запросов: Запросы, затрагивающие несколько шардов, могут быть сложными и менее эффективными.
// Управление данными: Требует управления распределением данных между шардами и мониторинга их состояния.
// Потенциальные проблемы с балансировкой: Неравномерное распределение данных может привести к перегрузке некоторых шардов.
package main

import (
	"fmt"
	"hash/fnv"
)

// Shard - структура, представляющая шард
type Shard struct {
	ID   int
	Data map[string]string
}

// NewShard - создает новый шард
func NewShard(id int) *Shard {
	return &Shard{
		ID:   id,
		Data: make(map[string]string),
	}
}

// ShardManager - структура для управления шардированием
type ShardManager struct {
	shards []*Shard
}

// NewShardManager - создает новый менеджер шардов
func NewShardManager(numShards int) *ShardManager {
	shards := make([]*Shard, numShards)
	for i := 0; i < numShards; i++ {
		shards[i] = NewShard(i)
	}
	return &ShardManager{shards: shards}
}

// getShard - определяет, в какой шард поместить данные
func (sm *ShardManager) getShard(key string) *Shard {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	shardIndex := int(hash.Sum32()) % len(sm.shards)
	return sm.shards[shardIndex]
}

// Set - добавляет данные в соответствующий шард
func (sm *ShardManager) Set(key, value string) {
	shard := sm.getShard(key)
	shard.Data[key] = value
	fmt.Printf("Key '%s' stored in shard %d\n", key, shard.ID)
}

// Get - извлекает данные из соответствующего шарда
func (sm *ShardManager) Get(key string) (string, bool) {
	shard := sm.getShard(key)
	value, exists := shard.Data[key]
	return value, exists
}

func main() {
	shardManager := NewShardManager(3)

	// Добавляем данные
	shardManager.Set("user1", "Alice")
	shardManager.Set("user2", "Bob")
	shardManager.Set("user3", "Charlie")

	// Извлекаем данные
	if value, exists := shardManager.Get("user1"); exists {
		fmt.Printf("Retrieved value for 'user1': %s\n", value)
	} else {
		fmt.Println("Key 'user1' not found")
	}
}
