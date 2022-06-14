package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Официант принимает заказ от клиента.
//Официант отдает заказ какому-то повару.
//Шеф-повар готовит заказ.
//Шеф-повар отдает приготовленное блюдо какому-нибудь официанту (не обязательно тому же официанту, который принимал заказ).
//Официант подает блюдо клиенту.

var wg sync.WaitGroup

type Order struct {
	id int
}

func getWaiter() string {
	waiters := []string{
		"waiter 1",
		"waiter 2",
		"waiter 3",
	}

	idx := rand.Intn(len(waiters))

	return waiters[idx]
}

func getChef() string {
	chefs := []string{
		"chef 1",
		"chef 2",
		"chef 3",
		"chef 4",
	}

	idx := rand.Intn(len(chefs))

	return chefs[idx]
}

func takeOrder(id int, orderChan chan<- Order) {
	defer wg.Done()
	waiter := getWaiter()

	fmt.Printf("%s has taken order number %v\n", waiter, id)

	orderChan <- Order{
		id: id,
	}
}

func cookOrder(orderChan <-chan Order, deliveryChan chan<- Order) {
	defer wg.Done()
	chef := getChef()

	order := <-orderChan
	fmt.Printf("%s is cooking order number %v\n", chef, order.id)

	deliveryChan <- order
}

func bringOrder(deliveryChan <-chan Order) {
	defer wg.Done()
	waiter := getWaiter()
	order := <-deliveryChan

	fmt.Printf("%s has brought dishes for order %v\n", waiter, order.id)
}

func main() {
	orderChan := make(chan Order)
	deliveryChan := make(chan Order)

	for orderId := 0; orderId < 5; orderId++ {
		wg.Add(3)
		go takeOrder(orderId, orderChan)
		go cookOrder(orderChan, deliveryChan)
		go bringOrder(deliveryChan)

		wg.Wait()
	}
}
