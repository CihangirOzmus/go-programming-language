package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	lock  sync.Mutex
}

func increaseCounter(counter *Counter) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
}

func increaseCounterWithChannel(counter *Counter, ch chan bool) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
	ch <- true
}

func increaseCounterWithWaitGroup(counter *Counter, wg *sync.WaitGroup) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	wg.Done()
	fmt.Println(counter.value)
}

func main() {
	counter := Counter{}
	/*time.Sleep to wait
	for i := 0; i < 100; i++ {
		go increaseCounter(&counter)
	}
	time.Sleep(2 * time.Second)*/

	/*channel receiver to wait
	ch := make(chan bool)
	for i := 0; i < 100; i++ {
		go increaseCounterWithChannel(&counter, ch)
	}
	for i := 0; i < 100; i++ {
		<-ch
	}*/

	//wait group to wait
	//wg := sync.WaitGroup{}
	//wg.Add(100)
	//
	//for i := 0; i < 100; i++ {
	//	go increaseCounterWithWaitGroup(&counter, &wg)
	//}
	//wg.Wait()

	// wait group to wait inside the loop
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increaseCounterWithWaitGroup(&counter, &wg)
	}
	wg.Wait()

	fmt.Println("Done")
}
