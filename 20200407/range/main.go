package main

import (
    "fmt"
)

func main() {
    rangelist()
}

func rangelist() {
    characters := []byte(string("123456789"))
    
    i := 1
    for _, c := range characters[i:] {
        i+=1
        fmt.Print("", string(c)) 
    }
    fmt.Println("")
}
