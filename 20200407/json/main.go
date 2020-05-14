package main

import (
    "log"
    "encoding/json"
)


func main() {
    input := "{\"name\": \"changchun\", \"score\": {\"tiyu\": 88, \"yuwen\":44, \"yishu\": \"weizhi\"}}"   

    var info Info
    if err := json.Unmarshal([]byte(input), &info); err != nil {
        log.Println("unmarshel failed, err:%s", err)
        return
    }
 
    log.Println("finish, info:", info)
    log.Println("info.Score.size:", len(info.Score))
}


type Info struct {
    Name string `json:"name"`
    Success bool `json:"success"`
    Score map[string]interface{} `json:"score"`
}

type Scores struct {
    Paras map[string]interface{}
}
