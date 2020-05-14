package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := sync.WaitGroup{}

	for i := 0; i < 300; i++ {
		done.Add(1)
		go worcker(i, &done)
	}
	done.Wait()
}

func worcker(i int, done *sync.WaitGroup) {
	defer done.Done()

	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	count := 1
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick ", i, " at", t.Format("2006-01-02_15:04:05.000"))
			count++
			if count > 1000 {
				return
			}
		}
	}
}
