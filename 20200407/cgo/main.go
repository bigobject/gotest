package main

import (
	"fmt"
	"os"
	"test/cgo/hello"
	"test/cgo/wavmatch"
)

func main() {
	ret := hello.SayHello("hello world\n")
	fmt.Println("ret:", ret)

	var matcher wavmatch.ZyCosWavMatch
	if err := matcher.ReloadTemplate("./wav/"); err != nil {
		fmt.Println("err:", err)
		return
	}

	voice, err := getVoice("./rings/yhzm_5s.wav")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	reason, err := matcher.Match("e2323232", voice)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("reason:", reason)
}

func getVoice(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("getVoice:%s", err)
	}
	defer file.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("getVoice:%s", err)
	}

	if fileInfo.Size() < 44 {
		return "", fmt.Errorf("fileInfo illegal")
	}

	Voice := make([]byte, fileInfo.Size())
	if n, err := file.Read(Voice); err != nil || int64(n) != fileInfo.Size() {
		return "", fmt.Errorf("getVoice:%s", err)
	}

	return string(Voice[44:]), nil
}
