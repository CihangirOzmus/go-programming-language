package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	someOps()              // 1
	double(1)              // 2
	double(3)              // 2
	double(5)              // 2
	fmt.Println(triple(4)) // 3

	filenames := []string{"a", "b", "c", "d", "e", "f", "g"} // 4
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			log.Println("All files processed")
		}
	}
}

// 1. track start end time with defer
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

// 2. logging with defer
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// 3. using return value within defer
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

// 4. use defer in functions otherwise risky; could run out of file descriptors
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...
	return nil
}
