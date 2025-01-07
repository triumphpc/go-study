// Паттерн "Scheduler Agent Supervisor" используется для управления выполнением задач в распределенных системах.
//Он состоит из трех основных компонентов: планировщика (Scheduler), агента (Agent) и супервайзера (Supervisor).
//Этот паттерн помогает координировать выполнение задач, обеспечивать их надежность и управлять ресурсами.
//
//Компоненты паттерна
//Scheduler (Планировщик): Отвечает за планирование задач и распределение их между агентами. Он определяет, какие задачи должны быть выполнены и когда.
//
//Agent (Агент): Выполняет задачи, назначенные планировщиком. Агенты могут быть распределены по различным узлам в системе и отвечают за выполнение конкретных задач.
//
//Supervisor (Супервайзер): Мониторит выполнение задач и состояние агентов. Он может перезапускать задачи или агентов в случае сбоев и обеспечивает надежность системы.
//
//Когда использовать Scheduler Agent Supervisor Pattern
//Распределенные системы: Когда задачи должны выполняться на нескольких узлах или серверах.
//Управление ресурсами: Когда необходимо эффективно распределять задачи между доступными ресурсами.
//Обеспечение надежности: Когда требуется автоматическое восстановление задач или агентов в случае сбоев.
//Планирование задач: Когда задачи должны выполняться в определенное время или с определенной периодичностью.

//Плюсы Scheduler Agent Supervisor Pattern
//Масштабируемость: Позволяет легко добавлять новые агенты для увеличения производительности.
//Надежность: Обеспечивает автоматическое восстановление задач и агентов в случае сбоев.
//Гибкость: Позволяет легко изменять планирование задач и управление ресурсами.
//Централизованное управление: Обеспечивает централизованное управление задачами и ресурсами.

// Минусы Scheduler Agent Supervisor Pattern
// Сложность реализации: Требует тщательной реализации и управления взаимодействием между компонентами.
// Задержки: Может возникнуть задержка в выполнении задач из-за распределения и мониторинга.
// Единая точка отказа: Планировщик может стать единой точкой отказа, если не предусмотрены механизмы отказоустойчивости.
// Зависимость от сети: Зависит от сетевой доступности и может быть подвержен сетевым сбоям.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task - структура, представляющая задачу
type Task struct {
	ID int
}

// Scheduler - планировщик задач
type Scheduler struct {
	tasks chan Task
}

// NewScheduler - создает новый планировщик
func NewScheduler() *Scheduler {
	return &Scheduler{tasks: make(chan Task, 10)}
}

// ScheduleTask - добавляет задачу в очередь
func (s *Scheduler) ScheduleTask(task Task) {
	s.tasks <- task
}

// Close - закрывает канал задач
func (s *Scheduler) Close() {
	close(s.tasks)
}

// Agent - агент, выполняющий задачи
type Agent struct {
	ID int
}

// Work - выполняет задачи из очереди
func (a *Agent) Work(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Agent %d processing task %d\n", a.ID, task.ID)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Симулируем выполнение задачи
	}
}

// Supervisor - супервайзер, следящий за выполнением задач
type Supervisor struct {
	scheduler *Scheduler
	agents    []*Agent
}

// NewSupervisor - создает нового супервайзера
func NewSupervisor(scheduler *Scheduler, agents []*Agent) *Supervisor {
	return &Supervisor{scheduler: scheduler, agents: agents}
}

// Monitor - следит за выполнением задач
func (s *Supervisor) Monitor() {
	var wg sync.WaitGroup
	for _, agent := range s.agents {
		wg.Add(1)
		go agent.Work(s.scheduler.tasks, &wg)
	}
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	scheduler := NewScheduler()
	agents := []*Agent{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	supervisor := NewSupervisor(scheduler, agents)

	// Планируем задачи
	for i := 1; i <= 10; i++ {
		scheduler.ScheduleTask(Task{ID: i})
	}

	// Закрываем канал задач после добавления всех задач
	scheduler.Close()

	// Запускаем мониторинг и выполнение задач
	supervisor.Monitor()
}
