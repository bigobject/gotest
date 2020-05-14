package main

import "fmt"

func main() {
	proc(1)
	proc(1, 2)
	proc(1, 2, 3)
}

func proc(a ...int) {
	fmt.Println("count ", len(a), "a:", a)
}
