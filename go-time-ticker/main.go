package main

import (
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer LogElapsedTime("All operations", start)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); Run1() }()
	go func() { defer wg.Done(); Run2() }()
	go func() { defer wg.Done(); Run3() }()

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go TimeTicker(start, ticker, done)

	wg.Wait()    // wait for all runs to finish
	done <- true // stop the ticker goroutine
	ticker.Stop()
}
