package main

import (
	"fmt"
	"sync"
	"time"
)

var sessMap sync.Map

func main() {

	inputticker := time.NewTicker(20 * time.Millisecond)
	defer inputticker.Stop()

	outputticker := time.NewTicker(20 * time.Millisecond)
	defer outputticker.Stop()
	done := make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			case <-inputticker.C:
				if i < 10000 {
					data := make([]int, 1000)
					data[i%1000] = i
					sessMap.Store(i, data)
					fmt.Println("Store i:", i)
					i++

				}
			}
		}
	}()

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			case <-outputticker.C:
				val, ok := sessMap.Load(i)
				if !ok {
					break
				}
				data, ok := val.([]int)
				if !ok {
					break
				}
				data[i%1000] = i
				sessMap.Delete(i)
				fmt.Println("delete i:", i)
				i++
			}
		}
	}()

	time.Sleep(1600 * time.Second)

	done <- true
}
