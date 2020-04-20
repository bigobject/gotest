package main

import (
    "fmt"
)

func main() {
    test1()
    test2()
    test3()
}

func test1() {
	fmt.Println("first case")
        a := make([]int, 0, 10)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        b := append(a, 1)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        _ = append(a, 2)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        fmt.Println(b[0])  //结果为2
        // fmt.Println(a[0]) //编译失败
}

func test2() {
	fmt.Println("second case")
        a := make([]int, 0)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        b := append(a, 1)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        _ = append(a, 2)
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
        fmt.Println(b[0])  //结果为1
        // fmt.Println(a[0]) //编译失败

}

func test3() {
	fmt.Println("third case")
	a := []int{1, 2, 3, 4, 5}
	doAppend(a[0:2])
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
	fmt.Println(a)
}


func doAppend(a []int) {
        fmt.Println("len(a):", len(a), ", cap(a):", cap(a))
	_ = append(a, 0)
}
