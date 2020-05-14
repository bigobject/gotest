package main

import (
    "time"
    "fmt"
)


func main() {
    time.Sleep(time.Duration(30)*time.Second) 
}
var (
    currUnixTime int64
    currDateTime string
    currDateHour string
    currDateDay  string
)


func init() {
    now := time.Now()
    currUnixTime = now.Unix()
    currDateTime = now.Format("2006-01-02 15:04:05") + fmt.Sprintf(".%03d", now.Nanosecond()/1e6)
    currDateHour = now.Format("2006010215")
    currDateDay = now.Format("20060102")
    go func() {
        tm := time.NewTimer(time.Millisecond)
        if err := recover(); err != nil { // avoid timer panic
        }   
        for {
            now := time.Now()
            d := time.Millisecond
            tm.Reset(d)
            <-tm.C
            now = time.Now()
            currUnixTime = now.Unix()
            currDateTime = now.Format("2006-01-02 15:04:05") + fmt.Sprintf(".%03d", now.Nanosecond()/1e6)
            currDateHour = now.Format("2006010215")
            currDateDay = now.Format("20060102")
        }   
    }() 

    test()
}


func test(){
    i := 0
    for {
        select {
        case <-time.After(time.Second * time.Duration(2)):
            i++
            if i == 5{
                fmt.Println("break now")
                break 
            }
            fmt.Println("inside the select: ")
        }
        fmt.Println("inside the for: ")
    }
}

