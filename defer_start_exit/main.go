package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	someOps()
	double(1)
	double(3)
	double(5)
}

// track start end time with defer
func someOps() {
	defer traceTime("bigSlowOpMsg")()
	time.Sleep(3 * time.Second)
}

func traceTime(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// logging with defer
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
