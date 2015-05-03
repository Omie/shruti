package main

import (
	"log"
	"os"
)

var responseQueue *TTSResponseQueue

func main() {

	responseQueue = NewTTSResponseQueue()

	accessKey := os.Getenv("SHRUTI_ACCESSKEY")
	secretKey := os.Getenv("SHRUTI_SECRETKEY")
	if accessKey == "" || secretKey == "" {
		log.Println("main: environment variables not set")
		return
	}

	InitIvona(accessKey, secretKey)

	go RunResponseWorker()

	if err := StartHTTPServer("127.0.0.1", "9574"); err != nil {
		log.Println("Error starting server", err)
	}

}
