package main

import (
    "regexp"
    "fmt"
)


func main() {
    fmt.Println(regexp.Match("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{1,5}", []byte("a0.23.23.23:232")))
}
