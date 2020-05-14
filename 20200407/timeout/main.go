package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	input := make(chan int, 3)

    var wg sync.WaitGroup
    wg.Add(1)

	go func() {
        fmt.Println("begin")
        defer wg.Done()

		timeouted := false
		stopinput := false
		timeout := time.After(time.Millisecond * 1000)
		for {
			select {
			case _, ok := <-input:
				if !ok {
					stopinput = true
					fmt.Println("stopinput")
					break
				}

			case <-timeout:
				timeouted = true
				fmt.Println("timeout")
				break
			}
			if timeouted && stopinput {
				break
			}
		}
        
        fmt.Println("end")
	}()

    go  func() {
        var (i =0)

        for{
            i++
            input <- i
            time.Sleep(50 * time.Millisecond)
            if i == 30 {
                close(input)
                break
            }
        } 
    }()
    wg.Wait()

}
