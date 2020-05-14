package main

import (
    "log"
)


func main() {
    var i int
    i = 2
    f(&i)
    log.Println(i)
    p(&i)
    log.Println(i)

    j := &i
    pp(&j)
    
    log.Println(*j)

}

func f(i *int) {
    *i = 3
}
func p(i *int) {
    var j int
    j =4
    i = &j
}

func pp(i **int) {
    var j int
    j = 5

    var k *int
    k = &j
    *i = k
}


