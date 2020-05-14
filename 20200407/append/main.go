package main

import (
	"fmt"
	"sync"
)

var input []byte
var mutex sync.Mutex

func main() {
	done := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		done.Add(1)
		go worker(i, &done)
	}
	done.Wait()
	fmt.Println("len input is:", len(input), ", cap is:", cap(input))
}

func worker(id int, done *sync.WaitGroup) {
	for len(input) < 10000000 {
		//mutex.Lock()
		input = append(input, []byte(string("sdsdfs"))...)
		fmt.Println("worker", id, " set input len:", len(input), ", cap:", cap(input))
		//mutex.Unlock()
	}

	done.Done()
}
