package main

import (
    "log"
    "fmt"
)


func main() {
    err := fmt.Errorf("sdsdwe")
    tst := fmt.Errorf("tst:%s", err)

    log.Println(tst)
}


