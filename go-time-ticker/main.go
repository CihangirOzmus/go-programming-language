package main

import (
	"fmt"
	"sync"
	"time"
)

const TimePrecision int = 3
const ShowElapsedTime bool = true

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() { defer wg.Done(); run1() }()
	go func() { defer wg.Done(); run2() }()
	go func() { defer wg.Done(); run3() }()

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go TimeTicker(start, ticker, done)

	wg.Wait()    // wait for all runs to finish
	done <- true // stop the ticker goroutine
	ticker.Stop()

	LogElapsedTime("All operations", start)
}

func run1() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 1", time.Now())
	}
	time.Sleep(1 * time.Second)
	fmt.Println("run1 done")
}

func run2() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 1", time.Now())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("run2 done")
}

func run3() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 3", time.Now())
	}
	time.Sleep(3 * time.Second)
	fmt.Println("run3 done")
}
