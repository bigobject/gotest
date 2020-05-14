package main

import (
    "fmt"
)

type student struct {
    Name string
    Age  int
}

func main() {
    m := make(map[string] student)
    stus := []student{
        {Name: "zhao1", Age: 12},
        {Name: "zhao2", Age: 13},
        {Name: "zhao3", Age: 14},
    }
    for _, stu := range stus {
        stu.Age = stu.Age+10
        fmt.Println(stu.Name, stu.Age)
    }
    for _, v := range stus {
        fmt.Println(v.Name, v.Age)
    }

    fmt.Println("===========")
    for i := 0; i < len(stus); i++ {
        m[stus[i].Name] = stus[i]
    }
    for _, v := range m {
        fmt.Println(v.Name, v.Age)
    }
}
