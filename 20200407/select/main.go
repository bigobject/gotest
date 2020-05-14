package main

import (
    "time"
    "log"
)


func main() {

    timer := time.NewTicker(time.Duration(1) * time.Second)

    for {
        select {
            case <- timer.C:
                log.Println("timout")
        }
    }

}


