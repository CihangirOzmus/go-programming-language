package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOp()
}

func bigSlowOp() {
	defer trace("bigSlowOpMsg")()
	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
