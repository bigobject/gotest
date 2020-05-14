package main

import (
    "log"
    "net/http"
    "time"
    "fmt"
)


func main() {
	InitServer("127.0.0.1", "23233","/first/")
	InitServer("127.0.0.1", "23233","/second/")

	time.Sleep(time.Second)
}

func onComposeResult(w http.ResponseWriter, r *http.Request) {

}
func InitServer(ip, port, path string) {
	go func() {
		http.HandleFunc(path, onComposeResult)
		g_ListenResultAddr := "http://" + ip + ":" + port+path
		log.Println("init server, addr:", g_ListenResultAddr)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Println(fmt.Sprintf("Tts InitServer failed, err:%s", err))
		}
	}()
}
