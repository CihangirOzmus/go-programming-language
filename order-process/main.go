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

const count = 10

//var (
//	totalUpdates int
//	updateMutex  sync.Mutex
//)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	//orderChan := make(chan *Order) // unbuffered channel, synchronous
	orderChan := make(chan *Order, count) // buffered channel, asynchronous
	processedChan := make(chan *Order, count)

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(count) {
			orderChan <- order
			fmt.Printf("Order id: <%d>, status: <%s>: send to channel!\n", order.Id, order.Status)
		}
		close(orderChan) // best practise to close the channel
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case processedOrder, ok := <-processedChan:
				if !ok {
					fmt.Println("Processed channel closed.")
					return
				}
				fmt.Printf("Processed order id: <%d> with status: <%s>\n", processedOrder.Id, processedOrder.Status)
			case <-time.After(10 * time.Second):
				fmt.Println("Timed out waiting for processed order.")
			}
		}
	}()

	go processOrders(orderChan, processedChan, &wg)

	//for range 3 {
	//	go func() {
	//		defer wg.Done()
	//		for _, order := range orders {
	//
	//			updateOrderStatuses(order)
	//		}
	//	}()
	//}
	wg.Wait()

	//reportOrdersStatus(orders)
	fmt.Println("All orders are processed!")
	//fmt.Println("Total updates processed:", totalUpdates)
}

//func updateOrderStatuses(order *Order) {
//	order.mu.Lock()
//	time.Sleep(1 * time.Second)
//	order.Status = Statuses[rand.Intn(len(Statuses))]
//	fmt.Printf("Updating Order <%d> status is:  <%s>\n", order.Id, order.Status)
//	order.mu.Unlock()
//
//	updateMutex.Lock()
//	currentUpdates := totalUpdates
//	time.Sleep(5 * time.Millisecond)
//	totalUpdates = currentUpdates + 1
//	defer updateMutex.Unlock()
//}

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		time.Sleep(1 * time.Second)
		fmt.Printf("Processing order id: <%d>\n", order.Id)
		order.Status = Statuses[rand.Intn(len(Statuses))]
		outChan <- order
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := range count {
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
