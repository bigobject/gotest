package main

import (
    "log"
    "test/http/httputil"
)


func main() {
    resp, err :=  httputil.Request("www.baidu.com", "")
    if err != nil {
        log.Println("finish, err:", err)
        return
    }
    log.Println("finish, resp:", resp)
}

