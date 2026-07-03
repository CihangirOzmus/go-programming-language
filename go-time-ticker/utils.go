package main

import (
	"fmt"
	"time"
)

const TimePrecision int = 3

func LogElapsedTime(name string, start time.Time) {
	fmt.Printf("%s took: %.*f seconds\n", name, TimePrecision, time.Since(start).Seconds())
}

func TimeTicker(start time.Time, ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			elapsed := time.Since(start).Seconds()
			if elapsed == 1 {
				fmt.Printf("%.*f second elapsed\n", TimePrecision, elapsed)
			} else {
				fmt.Printf("%.*f seconds elapsed\n", TimePrecision, elapsed)
			}
		}
	}
}
