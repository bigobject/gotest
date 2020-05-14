package main

import (
    "strings"
    "log"
)


func main() {
    input := "test3.mp3"
    
    result := strings.Trim(input, ".mp3") + "_00.pcm"
    log.Println(result)
}


