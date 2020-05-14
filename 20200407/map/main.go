package main

import (
    "fmt"
    "time"
)


type session struct{
    a string
    b string
    c string
    d string
}
func timeCost(sid string) func() {
    start := time.Now()
    return func() {
        fmt.Println(sid, " timecost:", time.Since(start))
    }
}

func getlen(d map[string]*session, name string) int {
    defer timeCost(name)()
    return len(d)
}
func getlenint(d map[string]int, name string) int {
    defer timeCost(name)()
    return len(d)
}


var list = make(map[string]*session, 1)
var large = make(map[string]*session, 65535)
var simple = make(map[string]int, 65535)

func main() {
    fmt.Println("simple:", getlenint(simple, "simple"))
    fmt.Println("list:", getlen(list, "list"))
    fmt.Println("large:", getlen(large, "large"))
    fmt.Println("large:", getlen(large, "large"))
    fmt.Println("large:", getlen(large, "large"))
    fmt.Println("large:", getlen(large, "large"))
    fmt.Println("large:", getlen(large, "large"))
    fmt.Println("large:", getlen(large, "large"))
}

