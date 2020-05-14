package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go run(1, 3, &wg)
	go run(2, 100, &wg)

	wg.Wait()
}

func run(id, max int, wg *sync.WaitGroup) {
	for count := 0; count < max; count++ {
		time.Sleep(time.Second)
		fmt.Println("id:", id, ", count:", count)
	}
	panic(fmt.Sprintf("id:%d paniced", id))
}
