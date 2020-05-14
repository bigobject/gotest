package main

import (
    "log"
)


func main() {
    in := make(chan int, 10)
    in <- 3
    in <- 2
    close(in)
    for  v:= range in {
        log.Println("v:", v)
    }
}

