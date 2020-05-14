package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./test/1/sa"
    
    file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("saveWav failed, err:%s", err)
        return
    }
    defer file.Close()
    
    if _, err = file.Write([]byte("hello, ni shi bu shi sa")); err != nil {
        fmt.Println("write failed")
        return
    }

    fmt.Println("write success")
}
