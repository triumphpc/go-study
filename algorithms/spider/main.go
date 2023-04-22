package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	numTests       = 500   // Количество тестов в итерации
	currentCostDef = 4000  // Стоимость позиции
	costPercentDef = 1     // Изменение цены в процентах на шаге
	totalMoneyDef  = 20000 // Всего не счету
)

var percent float32 = 2 // Процент снятия
var maxSteps = 2        // Количество шагов

var currentCost float32
var lastCost float32

var costPercent float32
var totalMoney float32

var hasRate = false
var directionCount = make(map[string]int, 2)

var lastRate float32
var totalRate float32

func main() {
	var total float32 = 0
	for i := 0; i < 5; i++ {
		reset()
		run()
		total += totalMoney
	}

	fmt.Println("СРЕДНЕЕ: ", total/5)
	//fmt.Println("ВСЕГО: ", total)

}

func reset() {
	currentCost = currentCostDef
	lastCost = 0
	costPercent = costPercentDef
	totalMoney = totalMoneyDef
	hasRate = false
	lastRate = 0
	totalRate = 0
}

func run() {
	directionCount["up"] = 0
	directionCount["down"] = 0

	defer func() {
		if directionCount["up"] > 0 {
			sellUp()
		} else if directionCount["down"] > 0 {
			sellDown()
		}

	}()

	for i := 0; i < numTests; i++ {
		if !hasRate {
			setRate()

			continue
		}

		g, err := nextStep()
		if err != nil {
			continue
		}

		if g > 0 {
			if directionCount["up"] > 0 {
				if directionCount["up"] == maxSteps {
					sellUp()
				} else {
					rateUp(lastRate * 2)
				}
			} else {
				rateUp(lastRate * 2)
				lastRate = currentCost
				totalRate = currentCost
			}
		} else {
			if directionCount["down"] > 0 {
				if directionCount["down"] == maxSteps {
					sellDown()
				} else {
					rateDown(lastRate * 2)
				}
			} else {
				rateDown(lastRate * 2)
				lastRate = currentCost
				totalRate = currentCost
			}
		}
	}
}

func sellUp() {
	totalMoney += totalRate + (totalRate * percent / 100)
	fmt.Printf("SELL UP `%f`\n", totalMoney)

	hasRate = false
	directionCount["up"] = 0
	directionCount["down"] = 0
	lastRate = 0
	totalRate = 0
}

func sellDown() {
	totalMoney -= totalRate
	totalMoney += totalRate * percent / 100
	fmt.Printf("SELL DOWN`%f`\n", totalMoney)

	hasRate = false
	directionCount["up"] = 0
	directionCount["down"] = 0
	lastRate = 0
	totalRate = 0
}

func setRate() {
	r, _ := nextStep()

	if r > 0 {
		rateUp(currentCost)
	} else {
		rateDown(currentCost)
	}

}

func rateUp(rate float32) {
	directionCount["down"] = 0
	directionCount["up"]++
	totalMoney -= rate
	lastRate = rate
	totalRate += rate
	hasRate = true
	lastCost = currentCost

	fmt.Printf("Ставка up  `%f`, всего денег `%f`\n", rate, totalMoney)
}

func rateDown(rate float32) {
	directionCount["up"] = 0
	directionCount["down"]++
	totalMoney += rate
	totalRate += rate
	lastRate = rate
	hasRate = true
	lastCost = currentCost

	fmt.Printf("Ставка down.  Текущая стоимость `%f`, всего денег `%f`\n", currentCost, totalMoney)
}

func nextStep() (int, error) {
	// Получим текущую цену
	_ = getCurrentCost()
	fmt.Printf("Новая цена `%f`\n", currentCost)

	diff := 100 - (lastCost * 100 / currentCost)
	fmt.Printf("Отклонение `%f` процентов\n", diff)

	if diff > 0 && diff >= percent {
		fmt.Println("Верхний порог")

		return 1, nil
	}

	if diff < 0 && math.Abs(float64(diff)) >= float64(percent) {
		fmt.Println("Нижний порог")

		return 0, nil
	}

	return 0, errors.New("")
}

func getCurrentCost() float32 {
	rand.Seed(time.Now().UnixNano())
	p := (rand.Intn(100-1) + 1) % 2

	if p > 0 {
		currentCost += currentCost * costPercent / 100
	} else {
		currentCost -= currentCost * costPercent / 100
	}

	return currentCost
}

var ch int
var dir int

func getCurrentCost2() float32 {
	dir++
	if ch%2 == 0 {
		if dir == 2 {
			ch++
			dir = 0
		}
		currentCost += currentCost * costPercent / 100
	} else {
		if dir == 2 {
			ch++
			dir = 0
		}
		currentCost -= currentCost * costPercent / 100
	}

	return currentCost
}

func getCurrentCost3() float32 {
	if ch%2 == 0 {
		currentCost += currentCost * costPercent / 100
	} else {
		currentCost -= currentCost * costPercent / 100
	}

	return currentCost
}
