package main

import (
	"fmt"
	"time"
)

const ShowElapsedTime bool = true

func Run1() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 1", time.Now())
	}
	time.Sleep(1 * time.Second)
	fmt.Println("run1 done")
}

func Run2() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 1", time.Now())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("run2 done")
}

func Run3() {
	if ShowElapsedTime {
		defer LogElapsedTime("Run 3", time.Now())
	}
	time.Sleep(3 * time.Second)
	fmt.Println("run3 done")
}
