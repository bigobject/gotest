package main

import (
    "log"
    "runtime/pprof"
    "os"
    "flag"
    "sync"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

    var wait sync.WaitGroup
    wait.Add(1)
    
    go func(wait *sync.WaitGroup) {
        var sum int
        for i :=0; i< 10000009990; i++ {
            sum += i
        }
        log.Println("sum:", sum)
        wait.Done()
    } (&wait)

    wait.Wait()
    log.Println("finish")
}

