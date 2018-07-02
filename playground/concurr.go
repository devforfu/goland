package main

import (
	"math/rand"
	"time"
	"flag"
	"fmt"
)

func main() {
	start := time.Now()
	numPtr := flag.Int("num", 10, "number of concurrent processors")
	flag.Parse()

	ch := make(chan float64)

	for i := 0; i < *numPtr; i++ {
		go processor(ch)
	}

	for i := 0; i < *numPtr; i++ {
		duration := <-ch
		fmt.Printf("%02d processor duration: %.2fs\n", i, duration)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Total time elapsed: %.2fs\n", elapsed)
}

func processor(ch chan<-float64) {
	start := time.Now()
	time.Sleep(randomDelay())
	elapsed := time.Since(start).Seconds()
	ch <- elapsed
}

func randomDelay() time.Duration {
	r := rand.Intn(1000)
	return time.Duration(r) * time.Millisecond
}
