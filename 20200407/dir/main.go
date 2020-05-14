package main

import (
	"log"
	"os"
)

func main() {
	path := "./test/1/sa"

	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
