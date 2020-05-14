package main

import (
    "log"
    "encoding/json"
)


func main() {
    input := "{\"bot1\":\"http://video-cdn.wezhuiyi.com/3DMM0125/130300.mp4\",\"bot2\":\"http://video-cdn.wezhuiyi.com/blendshaps/test2/130300.mp4\",\"bot3\":\"http://video-cdn.wezhuiyi.com/blendshaps/test2/130300.mp4\",\"faqID\":0,\"text\":\"您好！我行客服热线进入人工服务是24小时服务的。如果是涉及公司业务咨询，工作时间是早上八点到晚上10点.\"}"

    var info Info
    if err := json.Unmarshal([]byte(input), &info); err != nil {
        log.Println("unmarshel failed, err:%s", err)
        return
    }
 
    log.Println("finish, info:", info)
}


type Info struct {
    Bot1  string `json:"bot1"`
    Bot2  string `json:"bot2"`
    Bot3  string `json:"bot3"`
    FaqId int `json:"faqID"`
    Text  string `json:"text"`
}

