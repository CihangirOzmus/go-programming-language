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

func main() {
	orders := generateOrders(10)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		processOrders(orders)
	}()
	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrdersStatus(orders)
	}()
	wg.Wait()

	fmt.Println("All orders are processed!")
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(1 * time.Second)
		order.Status = Statuses[rand.Intn(len(Statuses))]
		fmt.Printf("Updating Order <%d> status is:  <%s>\n", order.Id, order.Status)
	}
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
	for i := 0; i < 3; i++ {
		for _, order := range orders {
			time.Sleep(1 * time.Second)
			fmt.Printf("Reporting Order <%d> status is:  <%s>\n", order.Id, order.Status)
		}
	}
}
