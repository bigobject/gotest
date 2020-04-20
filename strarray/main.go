package main

import (
	"fmt"
)

var sds = []string{
	1: "1111",
	2: "2222",
}

func main() {
	fmt.Println("1:", sds[1])
	fmt.Println("2:", sds[2])
}
