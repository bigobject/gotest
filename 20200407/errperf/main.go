package main

import (
	"errors"
	"fmt"
	"time"
)

var myerr = errors.New("this is my error")

func timeCost(sid string) func() {
	start := time.Now()
	return func() {
		fmt.Println(sid, "delay:", time.Since(start))
	}
}
func main() {
	fmt.Println("begin")
	tststringerror()
	tstmyerr()
	fmt.Println("end")
}

func tststringerror() {
	err := errors.New("this is his error")
	defer timeCost("tststringerror")()
	for i := 0; i < 10000; i++ {
		if err == myerr {
			fmt.Println("nothing")
		}
	}
}

func tstmyerr() {
	defer timeCost("tstmyerr")()
	for i := 0; i < 10000; i++ {
		if myerr.Error() == "this is his error" {
			fmt.Println("nothing")
		}
	}
}
