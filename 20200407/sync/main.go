package main

import (
	"log"
    "sync"
)
var g_tst sync.Map
func main() {
    if _, ok := g_tst.Load(""); !ok{
        log.Println("load failed")
        return
    }

    log.Println("load success")
}
