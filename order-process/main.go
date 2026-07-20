package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Id     int
	Status OrderStatus
	mu     sync.Mutex
}

type OrderStatus string

var Statuses = []OrderStatus{
	"New",
	"Paid",
	"Shipped",
	"Delivered",
	"Cancelled",
	"Complete",
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	orders := generateOrders(10)

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {

				updateOrderStatuses(order)
			}
		}()
	}
	wg.Wait()

	reportOrdersStatus(orders)
	fmt.Println("All orders are processed!")
	fmt.Println("Total updates processed:", totalUpdates)
}

func updateOrderStatuses(order *Order) {
	order.mu.Lock()
	time.Sleep(1 * time.Second)
	order.Status = Statuses[rand.Intn(len(Statuses))]
	fmt.Printf("Updating Order <%d> status is:  <%s>\n", order.Id, order.Status)
	order.mu.Unlock()

	updateMutex.Lock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
	defer updateMutex.Unlock()
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(1 * time.Second)
		fmt.Printf("Processing order <%d>\n", order.Id)
		order.Status = Statuses[rand.Intn(len(Statuses))]
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{Id: i + 1, Status: Statuses[0]}
	}
	return orders
}

func reportOrdersStatus(orders []*Order) {
	for _, order := range orders {
		time.Sleep(1 * time.Second)
		fmt.Printf("Reporting Order <%d> status is:  <%s>\n", order.Id, order.Status)
	}
}
